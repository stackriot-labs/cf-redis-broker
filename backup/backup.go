package backup

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"

	"code.google.com/p/go-uuid/uuid"

	"github.com/pivotal-cf/cf-redis-broker/backup/s3bucket"
	"github.com/pivotal-cf/cf-redis-broker/backupconfig"
	"github.com/pivotal-cf/cf-redis-broker/log"
	"github.com/pivotal-cf/cf-redis-broker/redis/client"
	"github.com/pivotal-cf/cf-redis-broker/redisconf"
	"github.com/pivotal-golang/lager"
)

type Backup struct {
	Config *backupconfig.Config
}

// http://golang.org/pkg/time/#pkg-constants if you need to understand this crazy layout
const (
	filenameTimestampLayout = "200601021504"
	datePathLayout          = "2006/01/02"
)

func (backup Backup) Create(instancePath, dataSubDir, instanceID, planName string) error {

	log.Logger().Info("backup", lager.Data{
		"event":         "backup_create_starting",
		"instance_path": instancePath,
		"data_dir":      dataSubDir,
		"instance_id":   instanceID,
		"plan_name":     planName,
	})

	if err := backup.createSnapshot(instancePath); err != nil {
		log.Logger().Error("backup", err, lager.Data{
			"event":         "backup_create",
			"instance_path": instancePath,
		})
		return err
	}

	timestamp := time.Now().Format(filenameTimestampLayout)

	rdbFilePath := filepath.Join(instancePath, dataSubDir, "dump.rdb")
	if !fileExists(rdbFilePath) {
		log.Logger().Info("dump.rdb not found, skipping instance backup", lager.Data{
			"Local file": rdbFilePath,
		})
		return nil
	}

	// rename dump.rdb
	tempRdbPath := filepath.Join(instancePath, dataSubDir, uuid.New())
	if err := os.Rename(rdbFilePath, tempRdbPath); err != nil {
		log.Logger().Error("backup", err, lager.Data{
			"event":    "dump_rename",
			"old_path": rdbFilePath,
			"new_path": tempRdbPath,
		})
		return err
	}

	// move rdb back or delete
	defer cleanup(tempRdbPath, rdbFilePath)

	bucket, err := backup.getOrCreateBucket()
	if err != nil {
		log.Logger().Error("backup", err, lager.Data{
			"event": "get_or_create_bucket",
		})
		return err
	}

	err = backup.uploadToS3(instanceID, planName, tempRdbPath, timestamp, bucket)
	if err != nil {
		log.Logger().Error("backup", err, lager.Data{
			"event":      "upload_to_s3",
			"bucket":     bucket,
			"local_path": tempRdbPath,
			"timestamp":  timestamp,
			"plan_name":  planName,
			"instanceID": instanceID,
		})
		return err
	}

	log.Logger().Info("backup", lager.Data{
		"event": "backup_create_done",
	})

	return nil
}

func cleanup(tempRdbPath, rdbPath string) {
	if !fileExists(rdbPath) {
		if err := os.Rename(tempRdbPath, rdbPath); err != nil {
			log.Logger().Error("backup", err, lager.Data{
				"event":    "rdb_rename",
				"old_path": tempRdbPath,
				"new_path": rdbPath,
			})
		}
	}

	if fileExists(tempRdbPath) {
		if err := os.Remove(tempRdbPath); err != nil {
			log.Logger().Error("backup", err, lager.Data{
				"event": "remove_temp_rdb",
				"path":  tempRdbPath,
			})
		}
	}
}

func (backup Backup) getOrCreateBucket() (s3bucket.Bucket, error) {
	s3Client := s3bucket.NewClient(
		backup.Config.S3Configuration.EndpointUrl,
		backup.Config.S3Configuration.AccessKeyId,
		backup.Config.S3Configuration.SecretAccessKey,
	)

	return s3Client.GetOrCreate(backup.Config.S3Configuration.BucketName)
}

func (backup Backup) createSnapshot(instancePath string) error {
	instanceConf, err := redisconf.Load(path.Join(instancePath, "redis.conf"))
	if err != nil {
		log.Logger().Error("backup", err, lager.Data{
			"event": "backup_create_snapshot_load",
		})
		return err
	}

	client, err := client.Connect("localhost", instanceConf)
	if err != nil {
		log.Logger().Error("backup", err, lager.Data{
			"event": "backup_create_snapshot_connect",
		})
		return err
	}

	err = client.CreateSnapshot(backup.Config.BGSaveTimeoutSeconds)
	if err != nil {
		log.Logger().Error("backup", err, lager.Data{
			"event": "backup_create_snapshot_create_snapshot",
		})
		return err
	}
	return nil
}

func (backup Backup) uploadToS3(instanceID, planName, rdbFilePath string, timestamp string, bucket s3bucket.Bucket) error {
	log.Logger().Info("s3", lager.Data{
		"event": "uploading",
	})

	remotePath := fmt.Sprintf("%s/%s/%s_%s_%s_redis_backup",
		backup.Config.S3Configuration.Path,
		time.Now().Format(datePathLayout),
		timestamp,
		instanceID,
		planName,
	)

	log.Logger().Info("Backing up instance", lager.Data{
		"Local file":  rdbFilePath,
		"Remote file": remotePath,
	})

	bucketPath := fmt.Sprintf("s3://%s%s", bucket.Name, remotePath)

	cmd := exec.Command(
		backup.Config.AwsCLIPath,
		"s3",
		"cp",
		rdbFilePath,
		bucketPath,
		"--endpoint-url",
		backup.Config.S3Configuration.EndpointUrl,
	)

	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", backup.Config.S3Configuration.AccessKeyId),
		fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", backup.Config.S3Configuration.SecretAccessKey),
	)

	log.Logger().Info(fmt.Sprintf("shelling out to aws cli to upload file %s to bucket %s", rdbFilePath, bucketPath))
	output, err := cmd.CombinedOutput()
	log.Logger().Info(string(output))

	log.Logger().Info("s3", lager.Data{
		"event": "uploading_done",
	})
	return err
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
