// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package secrets

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/gardener/gardener/pkg/utils/infodata"
)

// DataKeyVPNTLSAuth is the key in a secret data holding the vpn tlsauth key.
const DataKeyVPNTLSAuth = "vpn.tlsauth"

// VPNTLSAuthConfig contains the specification for a to-be-generated vpn tls authentication secret.
// The key will be generated by the provided VPNTLSAuthKeyGenerator. By default the openvpn command is used to generate the key if no generator function is specified.
type VPNTLSAuthConfig struct {
	Name                   string
	VPNTLSAuthKeyGenerator func() ([]byte, error)
}

// VPNTLSAuth contains the name and the generated vpn tls authentication key.
type VPNTLSAuth struct {
	Name       string
	TLSAuthKey []byte
}

// GetName returns the name of the secret.
func (s *VPNTLSAuthConfig) GetName() string {
	return s.Name
}

// Generate implements ConfigInterface.
func (s *VPNTLSAuthConfig) Generate() (DataInterface, error) {
	key, err := s.generateKey()
	if err != nil {
		return nil, err
	}

	return &VPNTLSAuth{
		Name:       s.Name,
		TLSAuthKey: key,
	}, nil
}

// GenerateInfoData implements ConfigInterface.
func (s *VPNTLSAuthConfig) GenerateInfoData() (infodata.InfoData, error) {
	key, err := s.generateKey()
	if err != nil {
		return nil, err
	}

	return NewPrivateKeyInfoData(key), nil
}

// GenerateFromInfoData implements ConfigInteface
func (s *VPNTLSAuthConfig) GenerateFromInfoData(infoData infodata.InfoData) (DataInterface, error) {
	data, ok := infoData.(*PrivateKeyInfoData)
	if !ok {
		return nil, fmt.Errorf("could not convert InfoData entry %s to PrivateKeyInfoData", s.Name)
	}

	return &VPNTLSAuth{
		Name:       s.Name,
		TLSAuthKey: data.PrivateKey,
	}, nil
}

// LoadFromSecretData implements infodata.Loader
func (s *VPNTLSAuthConfig) LoadFromSecretData(secretData map[string][]byte) (infodata.InfoData, error) {
	tlsAuthKey := secretData[DataKeyVPNTLSAuth]
	return NewPrivateKeyInfoData(tlsAuthKey), nil
}

func (s *VPNTLSAuthConfig) generateKey() (key []byte, err error) {
	if s.VPNTLSAuthKeyGenerator != nil {
		key, err = s.VPNTLSAuthKeyGenerator()
	} else {
		key, err = generateKeyDefault()
	}
	return
}

// SecretData computes the data map which can be used in a Kubernetes secret.
func (v *VPNTLSAuth) SecretData() map[string][]byte {
	data := map[string][]byte{
		DataKeyVPNTLSAuth: v.TLSAuthKey,
	}
	return data
}

func generateKeyDefault() ([]byte, error) {
	var (
		out bytes.Buffer
		cmd = exec.Command("openvpn", "--genkey", "--secret", "/dev/stdout")
	)

	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
