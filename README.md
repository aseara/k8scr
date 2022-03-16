

## 代码生成

```

ROOT_PACKAGE="github.com/aseara/k8scr/internal/api"
PACKAGE_VERSION="samplecrd:v1"

~/go/pkg/mod/k8s.io/code-generator@v0.23.4/generate-groups.sh all "internal/client" "$ROOT_PACKAGE" "$PACKAGE_VERSION" --go-header-file hack/custom-boilerplate.go.txt

mv -f $ROOT_PACKAGE/${PACKAGE_VERSION/://}/*.* internal/api/${PACKAGE_VERSION/://}/
rm -rf github.com

```