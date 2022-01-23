package gocapng

// CAPS from /usr/include/linux/capability.h
const (
	/*
		In a system with the [_POSIX_CHOWN_RESTRICTED] option defined, this
		overrides the restriction of changing file ownership and group ownership.
	*/
	CAPCHOWN Capability = iota

	/*
		Override all DAC access, including ACL execute access if [_POSIX_ACL] is
		defined. Excluding DAC access covered by CAP_LINUX_IMMUTABLE.
	*/
	CAPDACOverride

	/*
		Overrides all DAC restrictions regarding read and search on files and
		 directories, including ACL restrictions if [_POSIX_ACL] is defined.
		 Excluding DAC access covered by CAP_LINUX_IMMUTABLE.
	*/
	CAPDACReadSearch

	/*
		Overrides all restrictions about allowed operations on files, where file
		owner ID must be equal to the user ID, except where CAP_FSETID is applicable.
		It doesn't override MAC and DAC restrictions.
	*/
	CAPFOwner

	/*
		Overrides the following restrictions that the effective user ID shall match
		the file owner ID when setting the S_ISUID and S_ISGID bits on that file;
		that the effective group ID (or one of the supplementary group IDs) shall
		match the file owner ID when setting the S_ISGID bit on that file; that the
		S_ISUID and S_ISGID bits are cleared on successful return from
		chown(2) (not implemented).
	*/
	CAPFSetID

	/*
		Overrides the restriction that the real or effective user ID of a process
		sending a signal must match the real or effective user ID of the process
		receiving the signal.
	*/
	CAPKill

	/*
		Allows setgid(2) manipulation.
		Allows setgroups(2).
		Allows forged gids on socket credentials passing.
	*/
	CAPSetGID

	/*
	   Allows set*uid(2) manipulation (including fsuid).
	   Allows forged pids on socket credentials passing.
	*/
	CAPSetUID

	/**
	 ** Linux-specific capabilities
	 **/

	/*
		   Without VFS support for capabilities:
		     Transfer any capability in your permitted set to any pid,
		     remove any capability in your permitted set from any pid
		   With VFS support for capabilities (neither of above, but)
		     Add any capability from current's capability bounding set to the
				 current process' inheritable set.
		     Allow taking bits out of capability bounding set
		     Allow modification of the securebits for a process
	*/
	CAPSetPCap

	/*
		Allow modification of S_IMMUTABLE and S_APPEND file attributes
	*/
	CAPLinuxImmutable

	/*
		Allows binding to TCP/UDP sockets below 1024.
		Allows binding to ATM VCIs below 32.
	*/
	CAPNetBindService

	/*
	   Allow broadcasting, listen to multicast
	*/
	CAPNetBroadcast

	/*
		Allow interface configuration.
		Allow administration of IP firewall, masquerading and accounting.
		Allow setting debug option on sockets.
		Allow modification of routing tables.
		Allow setting arbitrary process / process group ownership on sockets.
		Allow binding to any address for transparent proxying (also via NET_RAW).
		Allow setting TOS (type of service).
		Allow setting promiscuous mode.
		Allow clearing driver statistics.
		Allow multicasting.
		Allow read/write of device-specific registers.
		Allow activation of ATM control sockets.
	*/
	CAPNetAdmin

	/*
		Allow use of RAW sockets.
		Allow use of PACKET sockets.
		Allow binding to any address for transparent proxying (also via NET_ADMIN).
	*/
	CAPNetRaw

	/*
		Allow locking of shared memory segments.
		Allow mlock and mlockall (which doesn't really have anything to do with IPC).
	*/
	CAPIPCLock

	/*
		Override IPC ownership checks.
	*/
	CAPIPCOwner

	/*
		Insert and remove kernel modules - modify kernel without limit.
	*/
	CAPSysModule

	/*
		Allow ioperm/iopl access.
		Allow sending USB messages to any device via /dev/bus/usb.
	*/
	CAPSysRawIO

	/*
	   Allow use of chroot().
	*/
	CAPSysChRoot

	/*
		Allow ptrace() of any process.
	*/
	CAPSysPTrace

	/*
		Allow configuration of process accounting.
	*/
	CAPSysPAcct

	/*
		Allow configuration of the secure attention key
		Allow administration of the random device
		Allow examination and configuration of disk quotas
		Allow setting the domainname
		Allow setting the hostname
		Allow calling bdflush()
		Allow mount() and umount(), setting up new smb connection
		Allow some autofs root ioctls
		Allow nfsservctl
		Allow VM86_REQUEST_IRQ
		Allow to read/write pci config on alpha
		Allow irix_prctl on mips (setstacksize)
		Allow flushing all cache on m68k (sys_cacheflush)
		Allow removing semaphores
		Used instead of CAP_CHOWN to "chown" IPC message queues, semaphores and shared memory
		Allow locking/unlocking of shared memory segment
		Allow turning swap on/off
		Allow forged pids on socket credentials passing
		Allow setting readahead and flushing buffers on block devices
		Allow setting geometry in floppy driver
		Allow turning DMA on/off in xd driver
		Allow administration of md devices (mostly the above, but some extra ioctls)
		Allow tuning the ide driver
		Allow access to the nvram device
		Allow administration of apm_bios, serial and bttv (TV) device
		Allow manufacturer commands in isdn CAPI support driver
		Allow reading non-standardized portions of pci configuration space
		Allow DDI debug ioctl on sbpcd driver
		Allow setting up serial ports
		Allow sending raw qic-117 commands
		Allow enabling/disabling tagged queuing on SCSI controllers and sending arbitrary SCSI commands
		Allow setting encryption key on loopback filesystem
		Allow setting zone reclaim policy
		Allow everything under CAP_BPF and CAP_PERFMON for backward compatibility.
	*/
	CAPSysAdmin

	/*
		Allow use of reboot().
	*/
	CAPSysBoot

	/*
		Allow raising priority and setting priority on other (different UID) processes
		Allow use of FIFO and round-robin (realtime) scheduling on own processes and
		  setting the scheduling algorithm used by another process.
		Allow setting cpu affinity on other processes
		Allow setting realtime ioprio class
		Allow setting ioprio class on other processes.
	*/
	CAPSysNice

	/*
		Override resource limits. Set resource limits.
		Override quota limits.
		Override reserved space on ext2 filesystem
		Modify data journaling mode on ext3 filesystem (uses journaling resources)

		NOTE: ext2 honors fsuid when checking for resource overrides, so you can override using fsuid too

		Override size restrictions on IPC message queues
		Allow more than 64hz interrupts from the real-time clock
		Override max number of consoles on console allocation
		Override max number of keymaps
		Control memory reclaim behavior
	*/
	CAPSysResource

	/*
		Allow manipulation of system cloc
		Allow irix_stime on mips
		Allow setting the real-time clock
	*/
	CAPSysTime

	/*
		Allow configuration of tty devices
		Allow vhangup() of tty
	*/
	CAPSysTTYConfig

	/*
		Allow the privileged aspects of mknod()
	*/
	CAPMkNod

	/*
		Allow taking of leases on files.
	*/
	CAPLease

	/*
		Allow writing the audit log via unicast netlink socket
	*/
	CAPAuditWrite

	/*
		Allow configuration of audit via unicast netlink socket
	*/
	CAPAuditControl

	/*
		Set or remove capabilities on files.
		Map uid=0 into a child user namespace.
	*/
	CAPSetFCap

	/*
		Override MAC access.
		The base kernel enforces no MAC policy.
		An LSM may enforce a MAC policy, and if it does and it chooses to implement
		capability based overrides of that policy, this is the capability it should use to do so.
	*/
	CAPMACOverride

	/*
		Allow MAC configuration or state changes.
		The base kernel requires no MAC configuration.
		An LSM may enforce a MAC policy, and if it does and it chooses to implement
		capability based checks on modifications to that policy or the data required
		to maintain it, this is the capability it should use to do so.
	*/
	CAPMACAdmin

	/*
		Allow configuring the kernel's syslog (printk behaviour)
	*/
	CAPSYSLOG

	/*
		Allow triggering something that will wake the system.
	*/
	CAPWakeAlarm

	/*
		Allow preventing system suspends
	*/
	CAPBlockSuspend

	/*
		Allow reading the audit log via multicast netlink socket
	*/
	CAPAuditRead

	/*
		Allow system performance and observability privileged operations
		using perf_events, i915_perf and other kernel subsystems
	*/
	CAPPerfmon

	/*
		CAP_BPF allows the following BPF operations:
		- Creating all types of BPF maps
		- Advanced verifier features
		  - Indirect variable access
		  - Bounded loops
		  - BPF to BPF function calls
		  - Scalar precision tracking
		  - Larger complexity limits
		  - Dead code elimination
		  - And potentially other features
		- Loading BPF Type Format (BTF) data
		- Retrieve xlated and JITed code of BPF programs
		- Use bpf_spin_lock() helper

		CAP_PERFMON relaxes the verifier checks further:
		- BPF progs can use of pointer-to-integer conversions
		- speculation attack hardening measures are bypassed
		- bpf_probe_read to read arbitrary kernel memory is allowed
		- bpf_trace_printk to print kernel memory is allowed

		CAP_SYS_ADMIN is required to use bpf_probe_write_user.

		CAP_SYS_ADMIN is required to iterate system wide loaded
		programs, maps, links, BTFs and convert their IDs to file descriptors.

		CAP_PERFMON and CAP_BPF are required to load tracing programs.
		CAP_NET_ADMIN and CAP_BPF are required to load networking programs.
	*/
	CAPBPF

	/*
		Allow checkpoint/restore related operations
		Allow PID selection during clone3()
		Allow writing to ns_last_pid
	*/
	CAPCheckpointRestore
)
