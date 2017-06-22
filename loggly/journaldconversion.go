package loggly

import (
	"encoding/json"
	"time"

	"github.com/uswitch/journald-forwarder/journald"
)

func ProcessJournal(c <-chan journald.Entry, errC chan<- error, uri string) {
	for msg := range c {

		logglyEntry := JournalEntry{
			PID:                     msg.PID,
			UID:                     msg.UID,
			GID:                     msg.GID,
			Comm:                    msg.Comm,
			Exe:                     msg.Exe,
			Cmdline:                 msg.Cmdline,
			CapEffective:            msg.CapEffective,
			AuditSession:            msg.AuditSession,
			AuditLoginID:            msg.AuditLoginID,
			SystemdGroup:            msg.SystemdGroup,
			SystemdSession:          msg.SystemdSession,
			SystemdUnit:             msg.SystemdUnit,
			SystemdUserInit:         msg.SystemdUserInit,
			SystemdOwnerUID:         msg.SystemdOwnerUID,
			SystemdSlice:            msg.SystemdSlice,
			SelinuxContext:          msg.SelinuxContext,
			SourceRealtimeTimestamp: msg.SourceRealtimeTimestamp,
			BootID:                  msg.BootID,
			MachineID:               msg.MachineID,
			Hostname:                msg.Hostname,
			Transport:               msg.Transport,
			Timestamp:               time.Unix(0, msg.RealtimeTimestamp*1000),
			MonotonicTimestamp:      msg.MonotonicTimestamp,
			CoredumpUnit:            msg.CoredumpUnit,
			CoredumpUserInit:        msg.CoredumpUserInit,
			ObjectPID:               msg.ObjectPID,
			ObjectUID:               msg.ObjectUID,
			ObjectGID:               msg.ObjectGID,
			ObjectComm:              msg.ObjectComm,
			ObjectExe:               msg.ObjectExe,
			ObjectCmdline:           msg.ObjectCmdline,
			ObjectAuditSession:      msg.ObjectAuditSession,
			ObjectAuditLoginID:      msg.ObjectAuditLoginID,
			ObjectSystemdCgroup:     msg.ObjectSystemdCgroup,
			ObjectSystemdSession:    msg.ObjectSystemdSession,
			ObjectSystemdUnit:       msg.ObjectSystemdUnit,
			ObjectSystemdUserInit:   msg.ObjectSystemdUserInit,
			ObjectSystemdOwnerUID:   msg.ObjectSystemdOwnerUID,
			Message:                 string(msg.Message),
			MessageID:               msg.MessageID,
			Priority:                msg.Priority,
			CodeFile:                msg.CodeFile,
			CodeLine:                msg.CodeLine,
			CodeFunc:                msg.CodeFunc,
			ErrNo:                   msg.ErrNo,
			SyslogFacility:          msg.SyslogFacility,
			SyslogIdentifier:        msg.SyslogIdentifier,
			ContainderID:            msg.ContainerID,
			ContainerFullID:         msg.ContainerFullID,
			ContainerName:           msg.ContainerName,
			ContainerTag:            msg.ContainerTag,
		}

		jsonEntry, err := json.Marshal(logglyEntry)
		if err != nil {
			errC <- err
			return
		}

		SendEvent(string(jsonEntry)[:], uri)
	}
}
