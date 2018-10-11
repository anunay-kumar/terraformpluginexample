## Go installation

wget https://dl.google.com/go/go1.11.1.linux-amd64.tar.gz

tar -C /usr/local -xzf go1.11.1.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

# import terraform package
go get github.com/hashicorp/terraform/helper/schema

# terraformpluginexample



## Create providers.go

## Create main.go for binaries



# Defining Resources

As a general convention, Terraform providers put each resource in their own file, named after the resource, prefixed with resource_. 
* To create an <provider-name>_server, this would be resource_server.go by convention.

## create resource_server.go

* The Create, Read, and Delete functions are required for a resource to be functional. 
* Terraform itself handles which function to call and with what data. 
* Based on the schema and current state of the resource, 
* Terraform can determine whether it needs to create a new resource, update an existing one, or destroy.

## Define CRUD functions in resource_server.go

# build with

$ go build -o terraform-provider-example  # Format:  go build -o terraform-provider-<provider-name>

## Implementing Destroy
 implement resourceServerDelete


# Data Sources
* Data sources are a special subset of resources which are read-only. 
* They are resolved earlier than regular resources and can be used as part of Terraform's interpolation.
* For data sources look at https://github.com/ashishrpandey/la-terraform-plugin-example


References: 
* https://golang.org/doc/install
* https://www.terraform.io/docs/extend/writing-custom-providers.html
* https://medium.com/@jozmo/creating-custom-terraform-providers-341311823fa2

