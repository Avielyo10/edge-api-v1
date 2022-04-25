package image

import (
	"encoding/json"
	"strings"

	"github.com/Avielyo10/edge-api/internal/common/models"
)

// Package represent a package installed on the image.
type Package struct {
	name string
}

// Packages are list of packages installed on the image.
type Packages struct {
	requiredPackages []Package // required packages, non removable
	packages         []Package
}

// NewPackages returns a new list of packages.
func NewPackages(packages ...string) Packages {
	requiredPackages := []Package{
		{"ansible"},
		{"rhc"},
		{"rhc-worker-playbook"},
		{"subscription-manager"},
		{"subscription-manager-plugin-ostree"},
		{"insights-client"},
	}
	// map required packages, avoid duplicates the required packages
	var requiredPackagesMap = make(map[string]bool)
	for _, pkg := range requiredPackages {
		requiredPackagesMap[pkg.name] = true
	}

	newPackages := make([]Package, 0, len(packages))
	for _, pkg := range packages {
		newPkg := NewPackage(pkg)
		if !newPkg.IsZero() && !requiredPackagesMap[newPkg.name] {
			newPackages = append(newPackages, newPkg)
		}
	}
	return Packages{requiredPackages: requiredPackages, packages: newPackages}
}

// NewPackage returns a new package.
func NewPackage(name string) Package {
	if !isValidPackage(name) {
		return Package{}
	}
	return Package{
		name: name,
	}
}

// isValidPackage returns true if the name is valid.
func isValidPackage(name string) bool {
	return name != "" && strings.TrimSpace(name) != ""
}

// IsZero returns true if the package is empty.
func (p Package) IsZero() bool {
	return p == Package{}
}

// Packages returns a pointer package array of packages + the required packages.
func (p Packages) Packages() []Package {
	var packages []Package
	packages = append(packages, p.requiredPackages...)
	packages = append(packages, p.packages...)
	return packages
}

// Add adds a package to the list.
func (p *Packages) Add(pkg ...Package) {
	for _, pkg := range pkg {
		if !pkg.IsZero() {
			p.packages = append(p.packages, pkg)
		}
	}
}

// Remove removes a package from the list.
func (p *Packages) Remove(pkgs ...Package) {
	// create a hashsets of the packages
	pkgsSet := make(map[string]bool)
	for _, pkg := range p.packages {
		pkgsSet[pkg.name] = true
	}
	// remove the packages
	for _, pkg := range pkgs {
		if pkgsSet[pkg.name] {
			delete(pkgsSet, pkg.name)
		}
	}
	// convert the map to a slice
	p.packages = make([]Package, 0, len(pkgsSet))
	for pkg := range pkgsSet {
		p.packages = append(p.packages, NewPackage(pkg))
	}
}

// AddPackage adds a package to the image.
func (i *Image) AddPackage(pkg ...Package) {
	if pkg != nil {
		i.packages.Add(pkg...)
	}
}

// RemovePackage removes a package from the image.
func (i *Image) RemovePackage(pkg ...Package) {
	if pkg != nil {
		i.packages.Remove(pkg...)
	}
}

// StringArray returns the packages as an array of strings.
func (p Packages) StringArray() []string {
	var packages []string
	for _, pkg := range p.Packages() {
		packages = append(packages, pkg.name)
	}
	return packages
}

// MarshalJSON creates a custom JSON marshaller.
func (p Packages) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.StringArray())
}

// UnmarshalJSON creates a custom JSON unmarshaller.
func (p *Packages) UnmarshalJSON(b []byte) error {
	if string(b) == `[]` {
		return nil
	}
	var packages []string
	if err := json.Unmarshal(b, &packages); err != nil {
		return err
	}
	*p = NewPackages(packages...)
	return nil
}

// ErrMissingRequiredPackage is raised when a required package is missing.
type ErrMissingRequiredPackage struct {
	pkg Package
}

// Error implements the error interface.
func (e ErrMissingRequiredPackage) Error() string {
	return "missing required package: " + e.pkg.name
}

// Has returns true if the package is present.
func (p Packages) Has(pkg Package) bool {
	for _, pkgIn := range p.Packages() {
		if pkgIn.name == pkg.name {
			return true
		}
	}
	return false
}

// MarshalGorm marshals the packages to a gorm model.
func (p Packages) MarshalGorm(account string) []models.Package {
	var packages []models.Package
	for _, pkg := range p.Packages() {
		packages = append(packages, models.Package{Account: account, Name: pkg.name})
	}
	return packages
}
