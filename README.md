# terraformpluginexample

#Create providers.go

#Create main.go for binaries

# build with
$ go build -o terraform-<provider-type>-<provider-name>

#Defining Resources

#As a general convention, Terraform providers put each resource in their own file, named after the resource, prefixed with resource_. 
# To create an <provider-name>_server, this would be resource_server.go by convention.
#create resource_server.go

# The Create, Read, and Delete functions are required for a resource to be functional. 
# There are other functions, but these are the only required ones. 
# Terraform itself handles which function to call and with what data. 
# Based on the schema and current state of the resource, 
# Terraform can determine whether it needs to create a new resource, update an existing one, or destroy.

## Define CRUD functions in resource_server.go



# Implementing Destroy
# implement resourceServerDelete


#Data Sources
# data sources are a special subset of resources which are read-only. 
# They are resolved earlier than regular resources and can be used as part of Terraform's interpolation.

