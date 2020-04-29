# This included makefile should define the 'custom' target rule which is called here.
# The (-) sign before `include` will do an implicit check if the file exists.
-include $(INCLUDE_MAKEFILE)

.PHONY: release
release: custom 
