package repository

import (
	"context"
	"register-service/model"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func InsertAuditLog(auditLog model.AuditLog) (string, error) {

	if auditLog.OldValue != auditLog.NewValue { /* Init log */
		_, file, _, _ := runtime.Caller(0)
		pc, _, _, _ := runtime.Caller(0)
		functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

		log := log.WithFields(log.Fields{
			"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
			"funciton":  functionName,
		})
		db := GetDb()
		ctx := context.Background()

		// Check if database is alive.
		err := db.PingContext(ctx)

		if err != nil {
			log.Errorf("#%v", err)
			return "NONE", err
		}

		stmt, err := db.Prepare(model.CREATE_SQL_AUDITLOG)

		if err != nil {
			log.Errorf(" %#v", err)
			return "ERROR", err
		}
		log.Info("auditLog:", auditLog)
		_, err = stmt.Exec(auditLog.MerchantId,
			auditLog.UserId,
			auditLog.ActionName,
			auditLog.EntityType,
			auditLog.EntityId,
			auditLog.PropName,
			auditLog.OldValue,
			auditLog.NewValue,
			auditLog.Remark,
			auditLog.MenuName,
			auditLog.ClientIpAddress,
			auditLog.CreatedAt)

		if err != nil {
			log.Errorf(" %#v", err)
			return "ERROR", err
		}
		log.Debugf("insert_channel_config ==> %v", "COMPLETE")
		return "COMPLETE", nil
	} else {
		return "NONE", nil
	}
}
