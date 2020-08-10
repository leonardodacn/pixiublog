package jobs

import (
	"sync"

	cron "pixiublog/crons"

	"github.com/astaxie/beego"
)

var (
	mainCron *cron.Cron
	workPool chan bool
	lock     sync.Mutex
)

func init() {
	if size, _ := beego.AppConfig.Int("jobs.pool"); size > 0 {
		workPool = make(chan bool, size)
	}
	mainCron = cron.New()
	mainCron.Start()
}

func AddJob(spec string, job *Job) bool {
	lock.Lock()
	defer lock.Unlock()

	if GetEntryById(job.JobKey) != nil {
		return false
	}
	err := mainCron.AddJob(spec, job)
	if err != nil {
		beego.Error("AddJob: ", err.Error())
		return false
	}
	//fmt.Println(job)
	return true
}

func RemoveJob(jobKey int) {
	mainCron.RemoveJob(func(e *cron.Entry) bool {
		if v, ok := e.Job.(*Job); ok {
			if v.JobKey == jobKey {
				return true
			}
		}
		return false
	})
}

func GetEntryById(jobKey int) *cron.Entry {
	entries := mainCron.Entries()
	for _, e := range entries {
		if v, ok := e.Job.(*Job); ok {
			if v.JobKey == jobKey {
				return e
			}
		}
	}
	return nil
}

func GetEntries(size int) []*cron.Entry {
	ret := mainCron.Entries()
	if len(ret) > size {
		return ret[:size]
	}
	return ret
}
