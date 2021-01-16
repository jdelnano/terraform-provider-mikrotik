package mikrotik

import (
	"fmt"
	"testing"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var originalBgpName string = "test-bgp-instance"
var originalAs string = "65532"
var originalRouterId string = "172.21.16.1"

func TestAccMikrotikBgpInstance_create(t *testing.T) {
	resourceName := "mikrotik_bgp_instance.bar"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMikrotikBgpInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBgpInstance(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccBgpInstanceExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", originalBgpName),
					resource.TestCheckResourceAttr(resourceName, "as", originalAs),
					resource.TestCheckResourceAttr(resourceName, "router_id", originalRouterId),
				),
			},
		},
	})
}

//func TestAccMikrotikBgpInstance_createAndPlanWithNonExistantBgpInstance(t *testing.T) {
//	resourceName := "mikrotik_pool.bar"
//	removeBgpInstance := func() {
//
//		c := client.NewClient(client.GetConfigFromEnv())
//		pool, err := c.FindBgpInstanceByName(originalBgpName)
//		if err != nil {
//			t.Fatalf("Error finding the pool by name: %s", err)
//		}
//		err = c.DeleteBgpInstance(pool.Id)
//		if err != nil {
//			t.Fatalf("Error removing the pool: %s", err)
//		}
//	}
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckMikrotikBgpInstanceDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccBgpInstance(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccBgpInstanceExists(resourceName),
//					resource.TestCheckResourceAttrSet(resourceName, "id")),
//			},
//			{
//				PreConfig:          removeBgpInstance,
//				Config:             testAccBgpInstance(),
//				ExpectNonEmptyPlan: false,
//			},
//		},
//	})
//}
//
//func TestAccMikrotikBgpInstance_updateBgpInstance(t *testing.T) {
//	resourceName := "mikrotik_pool.bar"
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckMikrotikBgpInstanceDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccBgpInstance(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccBgpInstanceExists(resourceName),
//					resource.TestCheckResourceAttr(resourceName, "name", originalBgpName),
//					resource.TestCheckResourceAttr(resourceName, "ranges", originalRanges),
//				),
//			},
//			{
//				Config: testAccBgpInstanceUpdatedName(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccBgpInstanceExists(resourceName),
//					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
//					resource.TestCheckResourceAttr(resourceName, "ranges", originalRanges),
//				),
//			},
//			{
//				Config: testAccBgpInstanceUpdatedRanges(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccBgpInstanceExists(resourceName),
//					resource.TestCheckResourceAttr(resourceName, "name", originalBgpName),
//					resource.TestCheckResourceAttr(resourceName, "ranges", updatedRanges),
//				),
//			},
//			{
//				Config: testAccBgpInstanceUpdatedComment(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccBgpInstanceExists(resourceName),
//					resource.TestCheckResourceAttr(resourceName, "name", originalBgpName),
//					resource.TestCheckResourceAttr(resourceName, "ranges", originalRanges),
//					resource.TestCheckResourceAttr(resourceName, "comment", updatedComment),
//				),
//			},
//		},
//	})
//}
//
//func TestAccMikrotikBgpInstance_import(t *testing.T) {
//	resourceName := "mikrotik_pool.bar"
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckMikrotikBgpInstanceDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccBgpInstance(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					testAccBgpInstanceExists(resourceName),
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

func testAccBgpInstance() string {
	return fmt.Sprintf(`
resource "mikrotik_bgp_instance" "bar" {
    name = "%s"
    as = 65532
    router_id = "%s"
}
`, originalBgpName, originalRouterId)
}

//func testAccBgpInstanceUpdatedName() string {
//	return fmt.Sprintf(`
//resource "mikrotik_pool" "bar" {
//    name = "%s"
//    ranges = "%s"
//}
//`, updatedName, originalRanges)
//}
//
//func testAccBgpInstanceUpdatedRanges() string {
//	return fmt.Sprintf(`
//resource "mikrotik_pool" "bar" {
//    name = "%s"
//    ranges = "%s"
//}
//`, originalBgpName, updatedRanges)
//}
//
//func testAccBgpInstanceUpdatedComment() string {
//	return fmt.Sprintf(`
//resource "mikrotik_pool" "bar" {
//    name = "%s"
//    ranges = "%s"
//    comment = "%s"
//}
//`, originalBgpName, originalRanges, updatedComment)
//}
//
func testAccBgpInstanceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("mikrotik_bgp_instance does not exist in the statefile")
		}

		c := client.NewClient(client.GetConfigFromEnv())

		bgpInstance, err := c.FindBgpInstance(rs.Primary.Attributes["name"])

		if err != nil {
			return fmt.Errorf("Unable to get the bgp instance with error: %v", err)
		}

		if bgpInstance == nil {
			return fmt.Errorf("Unable to get the bgp instance")
		}

		if bgpInstance.Name == rs.Primary.Attributes["name"] {
			return nil
		}
		return nil
	}
}

func testAccCheckMikrotikBgpInstanceDestroy(s *terraform.State) error {
	c := client.NewClient(client.GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_bgp_instance" {
			continue
		}

		bgpInstance, err := c.FindBgpInstance(rs.Primary.Attributes["name"])

		_, ok := err.(*client.NotFound)
		if !ok && err != nil {
			return err
		}

		if bgpInstance != nil {
			return fmt.Errorf("bgp instance (%s) still exists", bgpInstance.Name)
		}
	}
	return nil
}
