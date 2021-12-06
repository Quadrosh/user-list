# help
help:
	$(info  )
	$(info ******************************************************************* )
	$(info ****                      Maintenance                         ***** )
	$(info ****   Usage: make COMMAND                                    ***** )
	$(info ************************** Commands ******************************* )
	$(info *   run        - default app run )
	$(info *   initdb     - initialise database )
	$(info *   build-bin  - build bin file )
	$(info *   run-bin    - run bin file )
	$(info ******************************************************************* )

# ==============================================================================
# default run
run:
	go run cmd/api/main.go -dbhost=localhost -dbport=5432 -dbname=userlist -dbuser=postgres -dbpass=password  -port=8000 


# ==============================================================================
# initialise database
initdb:
	go run cmd/cli/main.go initdb -dbhost=localhost -dbport=5432 -dbname=userlist -dbuser=postgres -dbpass=password

# ==============================================================================
# build-bin
build-bin:
	go build -o build/api/bin  cmd/api/main.go

# ==============================================================================
# build
run-bin:
	build/api/bin   -dbhost=localhost -dbport=5432 -dbname=userlist -dbuser=postgres -dbpass=password  -port=8000 
