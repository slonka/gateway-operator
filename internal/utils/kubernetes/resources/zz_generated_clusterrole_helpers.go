// This file is generated by /hack/generators/kic-clusterrole-generator. DO NOT EDIT.

package resources

import (
	"fmt"

	"github.com/Masterminds/semver"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/kong/gateway-operator/internal/consts"
	"github.com/kong/gateway-operator/internal/utils/kubernetes/resources/clusterroles"
	"github.com/kong/gateway-operator/internal/versions"
)

// -----------------------------------------------------------------------------
// ClusterRole generator helper
// -----------------------------------------------------------------------------

// GenerateNewClusterRoleForControlPlane is a helper function that extract
// the version from the tag, and returns the ClusterRole with all the needed
// permissions.
func GenerateNewClusterRoleForControlPlane(controlplaneName string, image string) (*rbacv1.ClusterRole, error) {
	versionToUse := versions.DefaultControlPlaneVersion
	imageToUse := consts.DefaultControlPlaneImage
	var constraint *semver.Constraints

	if image != "" {
		v, err := versions.FromImage(image)
		if err != nil {
			return nil, err
		}
		supported, err := versions.IsControlPlaneImageVersionSupported(image)
		if err != nil {
			return nil, err
		}
		if supported {
			imageToUse = image
			versionToUse = v.String()
		}
	}

	semVersion, err := semver.NewVersion(versionToUse)
	if err != nil {
		return nil, err
	}

	constraint, err = semver.NewConstraint("<2.10, >=2.9")
	if err != nil {
		return nil, err
	}
	if constraint.Check(semVersion) {
		return clusterroles.GenerateNewClusterRoleForControlPlane_lt2_10_ge2_9(controlplaneName), nil
	}

	constraint, err = semver.NewConstraint(">=2.10")
	if err != nil {
		return nil, err
	}
	if constraint.Check(semVersion) {
		return clusterroles.GenerateNewClusterRoleForControlPlane_ge2_10(controlplaneName), nil
	}

	return nil, fmt.Errorf("version %s not supported", imageToUse)
}
