package applications

import (
	"net/http"

	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

type (
	//AppAggregate - aggregation object for application information
	AppAggregate struct {
		TotalInstanceCount int
		TotalAppCount      int
		JavaBPCount        int
		RubyBPCount        int
		NodeBPCount        int
		GOBPCount          int
		PythonBPCount      int
		PHPBPCount         int
		OtherBPCount       int
		StoppedStateCount  int
		StartedStateCount  int
		DiegoAppsCount     int
	}
	//AppSearch - search for apps object
	AppSearch struct {
		Client           cloudFoundryClient
		ClientTargetInfo *cf.CloudFoundryAPIInfo
		AppStats         *AppAggregate
	}
	cloudFoundryClient interface {
		QueryAPIInfo() (*cf.CloudFoundryAPIInfo, error)
		QueryUsers(int, int, string, string) (userList cf.UserAPIResponse, err error)
		Query(verb string, domain string, path string, args interface{}) (response *http.Response)
	}
)