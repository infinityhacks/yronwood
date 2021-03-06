package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/monzo/terrors"
)

var (
	ConfigListenAddr                = getConfigFromOSEnv("YRONWOOD_LISTEN_ADDR", ":8080", true)
	ConfigIndexRedirect             = getConfigFromOSEnv("YRONWOOD_INDEX_REDIRECT", "https://images.scy.email", true)
	ConfigStorageDirectoryPublic    = getConfigFromOSEnv("YRONWOOD_STORAGE_DIRECTORY_PUBLIC", "/images/uploads/public", true)
	ConfigStorageDirectoryUnlisted  = getConfigFromOSEnv("YRONWOOD_STORAGE_DIRECTORY_UNLISTED", "/images/uploads/big", true)
	ConfigStorageDirectoryPrivate   = getConfigFromOSEnv("YRONWOOD_STORAGE_DIRECTORY_PRIVATE", "/images/uploads/private", true)
	ConfigStorageDirectoryThumbnail = getConfigFromOSEnv("YRONWOOD_STORAGE_DIRECTORY_THUMBNAIL", "/images/uploads/thumbnail", true)
	ConfigAccessTypePublic          = getConfigFromOSEnv("YRONWOOD_ACCESS_TYPE_PUBLIC", "public", true)
	ConfigAccessTypeUnlisted        = getConfigFromOSEnv("YRONWOOD_ACCESS_TYPE_UNLISTED", "big", true)
	ConfigAccessTypePrivate         = getConfigFromOSEnv("YRONWOOD_ACCESS_TYPE_PRIVATE", "private", true)
	ConfigMaxFileSize               = getConfigFromOSEnv("YRONWOOD_MAX_FILE_SIZE", "25165824", true) // 24MB
	ConfigPermittedExtensions       = getConfigFromOSEnv("YRONWOOD_PERMITTED_EXTENSIONS", "jpeg|jpg|png|gif", true)
	ConfigAuthenticationSigningKey  = getConfigFromOSEnv("YRONWOOD_AUTHENTICATION_SIGHNING_KEY", "", false)
	ConfigAuthenticationBasicSecret = getConfigFromOSEnv("YRONWOOD_AUTHENTICATION_BASIC_SECRET", "", false)
	ConfigAuthenticationBasicSalt   = getConfigFromOSEnv("YRONWOOD_AUTHENTICATION_BASIC_SALT", "", false)
	ConfigCORSAllowedOrigin         = getConfigFromOSEnv("YRONWOOD_CORS_ALLOWED_ORIGIN", "https://images.scy.email", true)
)

// This is intended to run inside Kubernetes as a pod, so we just set service Configurations from deployment Configuration.
func getConfigFromOSEnv(key, defaultValue string, canBeEmpty bool) string {
	envValue := os.Getenv(key)
	if envValue != "" {
		return envValue
	}

	if !canBeEmpty {
		panic(terrors.InternalService("invalid_Config", fmt.Sprintf("Config value cannot be empty: %s", key), nil))
	}

	return defaultValue
}

// FileExtensionToContentType returns the appropriate HTTP content type for a given extension.
func FileExtensionToContentType(extension string) string {
	switch strings.ToLower(extension) {
	case "jpg", "jpeg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "gif":
		return "image/gif"
	}

	return "application/octet-stream"
}
