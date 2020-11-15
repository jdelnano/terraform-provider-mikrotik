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
var updatedComment string = "updated"
var updatedNextPool string = "spare"

// spare pool vars
var spareName string = updatedNextPool
var spareRanges string = "172.16.0.248-172.16.0.249"
var spareComment string = "spare-comment"
var spareNextPool string = "none"

func TestAccMikrotikPool_create(t *testing.T) {
	resourceName := "mikrotik_pool.bar"
	pool, _ := createSparePool()
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
					resource.TestCheckResourceAttr(resourceName, "nextpool", spareName),
				),
				//ExpectNonEmptyPlan: true,
			},
		},
	})
	destroySparePool(pool)
}

func TestAccMikrotikPool_updatePool(t *testing.T) {
	resourceName := "mikrotik_pool.bar"
	// create tmp pool needed to specify for 'next-pool' field in test resource
	//pool, _ := createSparePool()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMikrotikPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPool(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccPoolExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", originalName),
					resource.TestCheckResourceAttr(resourceName, "ranges", originalRanges),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccPoolUpdatedName(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccPoolExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "ranges", originalRanges),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccPoolUpdatedRanges(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccPoolExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", originalName),
					resource.TestCheckResourceAttr(resourceName, "ranges", updatedRanges),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccPoolUpdatedComment(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccPoolExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", originalName),
					resource.TestCheckResourceAttr(resourceName, "ranges", originalRanges),
					resource.TestCheckResourceAttr(resourceName, "comment", updatedComment),
				),
				ExpectNonEmptyPlan: true,
			},
			//{
			//	Config: testAccPoolUpdatedNextPool(),
			//	Check: resource.ComposeAggregateTestCheckFunc(
			//		testAccPoolExists(resourceName),
			//		resource.TestCheckResourceAttr(resourceName, "name", originalName),
			//		resource.TestCheckResourceAttr(resourceName, "ranges", originalRanges),
			//		resource.TestCheckResourceAttr(resourceName, "nextpool", updatedNextPool),
			//	),
			//	//ExpectNonEmptyPlan: true,
			//},
		},
	})
	//destroySparePool(pool)
}

func TestAccMikrotikPool_import(t *testing.T) {
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
					resource.TestCheckResourceAttrSet(resourceName, "id")),
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPool() string {
	return fmt.Sprintf(`
resource "mikrotik_pool" "bar" {
    name = "%s"
    ranges = "%s"
    nextpool = "%s"
}
`, originalName, originalRanges, spareName)
}

func testAccPoolUpdatedName() string {
	return fmt.Sprintf(`
resource "mikrotik_pool" "bar" {
    name = "%s"
    ranges = "%s"
}
`, updatedName, originalRanges)
}

func testAccPoolUpdatedRanges() string {
	return fmt.Sprintf(`
resource "mikrotik_pool" "bar" {
    name = "%s"
    ranges = "%s"
}
`, originalName, updatedRanges)
}

func testAccPoolUpdatedComment() string {
	return fmt.Sprintf(`
resource "mikrotik_pool" "bar" {
    name = "%s"
    ranges = "%s"
    comment = "%s"
}
`, originalName, originalRanges, updatedComment)
}

func testAccPoolUpdatedNextPool() string {
	return fmt.Sprintf(`
resource "mikrotik_pool" "bar" {
    name = "%s"
    ranges = "%s"
    nextpool = "%s"
}
`, originalName, originalRanges, updatedNextPool)
}

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

func createSparePool() (*client.Pool, error) {
	c := client.NewClient(client.GetConfigFromEnv())

	name := spareName
	ranges := spareRanges
	comment := spareComment
	nextpool := spareNextPool
	pool, err := c.AddPool(
		name,
		ranges,
		comment,
		nextpool,
	)

	if err != nil {
		return nil, err
	}

	return c.FindPool(pool.Id)
}

func destroySparePool(p *client.Pool) error {
	c := client.NewClient(client.GetConfigFromEnv())

	err := c.DeletePool(p.Id)

	return err
}
