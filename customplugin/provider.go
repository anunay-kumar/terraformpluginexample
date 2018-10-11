/*#root of the provider

#The *schema.Provider type describes the provider's properties including:

#the configuration keys it accepts
#the resources it supports
#any callbacks to configure
*/
package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
        return &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{
				"example_server": resourceServer(),
				},
        }
}
