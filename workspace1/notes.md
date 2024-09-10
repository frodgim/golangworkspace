* Run modules
    * cd workspace1\secondmodule
    * go run .
* Create a module
    * go mod init <modulename>
* Install dependencies:
    * go mod tidy

* Create a workspace 
    * go work init <./modulefoldername>
    * Edit a workspace. Open up go.work then add a new "use <./another_module_folder>"
