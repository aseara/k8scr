

## 代码生成

```

ROOT_PACKAGE="github.com/aseara/k8scr"
CUSTOM_RESOURCE_NAME="samplecrd"
CUSTOM_RESOURCE_VERSION="v1"

~/go/pkg/mod/k8s.io/code-generator@v0.23.4/generate-groups.sh all "pkg/client" "$ROOT_PACKAGE/pkg/apis" "$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION" --go-header-file hack/custom-boilerplate.go.txt

mv -f $ROOT_PACKAGE/pkg/apis/$CUSTOM_RESOURCE_NAME/$CUSTOM_RESOURCE_VERSION/*.* pkg/apis/$CUSTOM_RESOURCE_NAME/$CUSTOM_RESOURCE_VERSION/
rm -rf github.com

```