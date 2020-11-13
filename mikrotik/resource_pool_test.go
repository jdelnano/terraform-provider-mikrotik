package mikrotik

import (
	"fmt"
	"testing"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var originalName string = "test-pool"
var originalRanges string = "172.16.0.1-172.16.0.8,172.16.0.10"
var updatedName string = "test-pool-updated"
var updatedRanges string = "172.16.0.11-172.16.0.12"

func TestAccMikrotikPool_create(t *testing.T) {
	resourceName := "mikrotik_pool.bar"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMikrotikPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPool(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccPoolExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", originalName),
					resource.TestCheckResourceAttr(resourceName, "ranges", originalRanges),
					resource.TestCheckResourceAttr(resourceName, "nextpool", ""),
				),
			},
		},
	})
}

//func TestAccMikrotikPool_updateAddress(t *testing.T) {
//	resourceName := "mikrotik_pool.bar"
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckMikrotikPoolDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccPool(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccPoolExists(resourceName),
//					resource.TestCheckResourceAttr(resourceName, "address", originalName),
//					resource.TestCheckResourceAttr(resourceName, "macaddress", originalRanges),
//				),
//			},
//			{
//				Config: testAccPoolUpdatedIpAddress(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccPoolExists(resourceName),
//					resource.TestCheckResourceAttr(resourceName, "address", updatedName),
//					resource.TestCheckResourceAttr(resourceName, "macaddress", originalRanges),
//				),
//			},
//			{
//				Config: testAccPoolUpdatedMacAddress(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccPoolExists(resourceName),
//					resource.TestCheckResourceAttr(resourceName, "address", originalName),
//					resource.TestCheckResourceAttr(resourceName, "macaddress", updatedRanges),
//				),
//			},
//			{
//				Config: testAccPoolUpdatedBlockAccess(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccPoolExists(resourceName),
//					resource.TestCheckResourceAttr(resourceName, "address", originalName),
//					resource.TestCheckResourceAttr(resourceName, "macaddress", originalRanges),
//					resource.TestCheckResourceAttr(resourceName, "blocked", updatedBlockAccess),
//				),
//			},
//		},
//	})
//}
//
//func TestAccMikrotikPool_import(t *testing.T) {
//	resourceName := "mikrotik_pool.bar"
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckMikrotikPoolDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccPool(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccPoolExists(resourceName),
//					resource.TestCheckResourceAttrSet(resourceName, "id")),
//			},
//			{
//				ResourceName:      resourceName,
//				ImportState:       true,
//				ImportStateVerify: true,
//			},
//		},
//	})
//}
//
//func TestAccMikrotikPool_createDynamicDiff(t *testing.T) {
//	resourceName := "mikrotik_pool.bar"
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckMikrotikPoolDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccPoolDynamic(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccPoolExists(resourceName),
//					resource.TestCheckResourceAttrSet(resourceName, "id")),
//				ExpectNonEmptyPlan: true,
//			},
//		},
//	})
//}
//
func testAccPool() string {
	return fmt.Sprintf(`
resource "mikrotik_pool" "bar" {
    name = "%s"
    ranges = "%s"
    nextpool = ""
}
`, originalName, originalRanges)
}

//func testAccPoolDynamic() string {
//	return fmt.Sprintf(`
//resource "mikrotik_pool" "bar" {
//    comment = "bar"
//    address = "%s"
//    macaddress = "%s"
//    dynamic = true
//}
//`, originalName, originalRanges)
//}

//func testAccPoolUpdatedIpAddress() string {
//	return fmt.Sprintf(`
//resource "mikrotik_pool" "bar" {
//    comment = "bar"
//    address = "%s"
//    macaddress = "%s"
//}
//`, updatedName, originalRanges)
//}
//
//func testAccPoolUpdatedMacAddress() string {
//	return fmt.Sprintf(`
//resource "mikrotik_pool" "bar" {
//    comment = "bar"
//    address = "%s"
//    macaddress = "%s"
//}
//`, originalName, updatedRanges)
//}
//
//func testAccPoolUpdatedBlockAccess() string {
//	return fmt.Sprintf(`
//resource "mikrotik_pool" "bar" {
//    comment = "bar"
//    address = "%s"
//    macaddress = "%s"
//    blocked= "%s"
//}
//`, originalName, originalRanges, updatedBlockAccess)
//}

func testAccPoolExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("mikrotik_pool does not exist in the statefile")
		}

		c := client.NewClient(client.GetConfigFromEnv())

		pool, err := c.FindPool(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Unable to get the pool with error: %v", err)
		}

		if pool == nil {
			return fmt.Errorf("Unable to get the pool")
		}

		if pool.Id == rs.Primary.ID {
			return nil
		}
		return nil
	}
}

//func testAccCheckMikrotikPoolDestroyNow(resourceName string) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		rs, ok := s.RootModule().Resources[resourceName]
//		if !ok {
//			return fmt.Errorf("Not found: %s", resourceName)
//		}
//
//		if rs.Primary.ID == "" {
//			return fmt.Errorf("No pool lease Id is set")
//		}
//
//		c := client.NewClient(client.GetConfigFromEnv())
//
//		pool, err := c.FindPool(rs.Primary.ID)
//
//		_, ok = err.(*client.NotFound)
//		if !ok && err != nil {
//			return err
//		}
//
//		err = c.DeletePool(pool.Id)
//
//		if err != nil {
//			return err
//		}
//
//		return nil
//	}
//}

func testAccCheckMikrotikPoolDestroy(s *terraform.State) error {
	c := client.NewClient(client.GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_pool" {
			continue
		}

		pool, err := c.FindPool(rs.Primary.ID)

		_, ok := err.(*client.NotFound)
		if !ok && err != nil {
			return err
		}

		if pool != nil {
			return fmt.Errorf("pool (%s) still exists", pool.Id)
		}
	}
	return nil
}
