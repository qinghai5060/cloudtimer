package cmd

import (
	"context"
	"io/ioutil"
	"os"
	"src/CloudCron/pkg"

	"github.com/golang/glog"

	vultr "github.com/JamesClonk/vultr/lib"

	gcfg "gopkg.in/yaml.v2"

	"github.com/robfig/cron"
)

type CloudCron struct {
	Crons       *cron.Cron
	Context     context.Context
	CloudClient *vultr.Client

	Server *vultr.Server
}

func NewCloudCron(ctx context.Context, key string) *CloudCron {
	cc := &CloudCron{
		Crons:       cron.New(),
		Context:     ctx,
		CloudClient: vultr.NewClient(key, nil),
		Server:      nil,
	}

	return cc
}

func (c *CloudCron) Run(ctx context.Context, conf string) error {
	if _, err := os.Stat(conf); err != nil {
		return nil
	}
	job := &pkg.Job{
		StartTime:  "",
		StopTime:   "",
		Key:        "",
		RegionID:   0,
		PlanID:     0,
		OsID:       0,
		ServerName: "",
	}
	ymlFile, err := ioutil.ReadFile(conf)
	if err != nil {
		glog.Errorf("failed to read config . error : %s", err.Error())
		return err
	}

	if err := gcfg.Unmarshal(ymlFile, job); err != nil {
		return err
	}

	glog.Infof("config start time : %s, stop time : %s, servername: %s, plan id :%d, os id :%d ", job.StartTime, job.StopTime, job.ServerName, job.PlanID, job.OsID)

	creatServer := func() {
		s, err := c.CloudClient.CreateServer(job.ServerName, job.RegionID, job.PlanID, job.OsID, nil)
		if err != nil {
			glog.Errorf("failed to create server. error : %s", err.Error())
			return
		}

		c.Server = &s

		glog.Infof("create server succeed. instance is : %s, ip :%s ", s.ID, s.MainIP)
	}

	stopSever := func() {
		if err := c.CloudClient.DeleteServer(c.Server.ID); err != nil {
			glog.Errorf("failed to delete server . error : %s", err.Error())
			return
		}

		c.Server = nil
	}

	if err := c.Crons.AddFunc(job.StartTime, creatServer); err != nil {
		glog.Errorf("failed to add func for start cron. error : %s", err.Error())
		return err
	}

	if err := c.Crons.AddFunc(job.StopTime, stopSever); err != nil {
		glog.Errorf("failed to add stop cron . error : %s", err.Error())
		return err
	}

	c.Crons.Start()

	return nil
}
