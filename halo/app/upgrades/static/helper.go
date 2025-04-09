package static

import "github.com/omni-network/omni/lib/errors"

func LatestUpgrade() string {
	return UpgradeNames[len(UpgradeNames)-1]
}

// NextUpgrade returns the next upgrade name after the provided previous upgrade,
// or false if this the latest upgrade (no next), or an error if the name is not a
// valid upgrade.
func NextUpgrade(prev string) (string, bool, error) {
	if prev == "" { // Return the first upgrade
		return UpgradeNames[0], true, nil
	}

	for i, name := range UpgradeNames {
		if name != prev {
			continue
		}

		if i == len(UpgradeNames)-1 {
			return "", false, nil // No next upgrade
		}

		return UpgradeNames[i+1], true, nil
	}

	return "", false, errors.New("prev upgrade not found [BUG]")
}
