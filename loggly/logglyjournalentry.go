package loggly

import "time"

type JournalEntry struct {
	PID                     int       `json:"pid,omitempty"`
	UID                     int       `json:"uid,omitempty"`
	GID                     int       `json:"gid,omitempty"`
	Comm                    string    `json:"appName,omitempty"`
	Exe                     string    `json:"exe,omitempty"`
	Cmdline                 string    `json:"cmdline,omitempty"`
	CapEffective            string    `json:"capEffective,omitempty"`
	AuditSession            int       `json:"auditSession,omitempty"`
	AuditLoginID            string    `json:"auditLoginId,omitempty"`
	SystemdGroup            string    `json:"systemdCgroup,omitempty"`
	SystemdSession          string    `json:"systemdSession,omitempty"`
	SystemdUnit             string    `json:"systemdUnit,omitempty"`
	SystemdUserInit         string    `json:"systemdUserInit,omitempty"`
	SystemdOwnerUID         string    `json:"systemdOwnerUid,omitempty"`
	SystemdSlice            string    `json:"systemdSlice,omitempty"`
	SelinuxContext          string    `json:"selinuxContext,omitempty"`
	SourceRealtimeTimestamp int64     `json:"sourceRealtimeTimestamp,omitempty"`
	BootID                  string    `json:"bootId,omitempty"`
	MachineID               string    `json:"machineId,omitempty"`
	Hostname                string    `json:"hostname,omitempty"`
	Transport               string    `json:"transport,omitempty"`
	Timestamp               time.Time `json:"timestamp,omitempty"`
	MonotonicTimestamp      int64     `json:"monotonicTimestamp,omitempty"`
	CoredumpUnit            string    `json:"coredumpUnit,omitempty"`
	CoredumpUserInit        string    `json:"coredumpUserInit,omitempty"`
	ObjectPID               int       `json:"objectPid,omitempty"`
	ObjectUID               int       `json:"objectUid,omitempty"`
	ObjectGID               int       `json:"objectGid,omitempty"`
	ObjectComm              string    `json:"objectComm,omitempty"`
	ObjectExe               string    `json:"objectExe,omitempty"`
	ObjectCmdline           string    `json:"objectCmdline,omitempty"`
	ObjectAuditSession      string    `json:"objectAuditSession,omitempty"`
	ObjectAuditLoginID      string    `json:"objectAuditLoginId,omitempty"`
	ObjectSystemdCgroup     string    `json:"objectSystemdCgroup,omitempty"`
	ObjectSystemdSession    string    `json:"objectSystemdSession,omitempty"`
	ObjectSystemdUnit       string    `json:"objectSystemdUnit,omitempty"`
	ObjectSystemdUserInit   string    `json:"objectSystemdUserInit,omitempty"`
	ObjectSystemdOwnerUID   int       `json:"objectSystemdOwnerUid,omitempty"`
	Message                 string    `json:"message,omitempty"`
	MessageID               string    `json:"messageId,omitempty"`
	Priority                int       `json:"priority,omitempty"`
	CodeFile                string    `json:"codeFile,omitempty"`
	CodeLine                string    `json:"codeLine,omitempty"`
	CodeFunc                string    `json:"codeFunc,omitempty"`
	ErrNo                   int       `json:"errNo,omitempty"`
	SyslogFacility          string    `json:"facility,omitempty"`
	SyslogIdentifier        string    `json:"syslogIdentifier,omitempty"`
	ContainderID            string    `json:"containerId,omitempty"`
	ContainerFullID         string    `json:"containerFullId,omitempty"`
	ContainerName           string    `json:"containerName,omitempty"`
	ContainerTag            string    `json:"containerTag,omitempty"`
}
