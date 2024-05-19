package download

/*
file structure:

	instances
	|-instance
	|	|-instance specific artifact(s)
	|
	|-common
		|-textures
		|-artifacts
		|-jre

updater:
	- update all artifacts
	- update all required artifacts
	- check natives

natives:
	- there are 2 kinds, presumably lwgl 2 and 3
	- possible just have this in common dir
	- change natives dir depending on mc ver
*/