- decrypt a few sectors to verify the partition table, or provide a tool to try to guess the partition table or something
	- http://unix.stackexchange.com/questions/289389/what-are-the-differences-between-the-various-partition-tables
- delete outimages on error?
- don't show an error message on password.ErrInterrupted
- make a list of valid errors to call usage() on
- investigate firmware modules; will require direct ATA access
	- Linux: not libata, check smartctl and hdparm source (http://stackoverflow.com/questions/17366515/executing-ata-commands-on-hdd-from-user-space-application-using-libata)
	- Windows: DeviceIoControl(), just like the Unlock.exe tool
	- OS X: requires a kernel module? https://developer.apple.com/library/content/qa/qa1179/_index.html
	- others: not possible?

- Test
-- Test
