

## 代码生成

```

CUR_MOD="github.com/aseara/k8scr"
API_PACKAGE="$CUR_MOD/internal/api"
PACKAGE_VERSION="samplecrd:v1"

chmod a+x ~/go/pkg/mod/k8s.io/code-generator@v0.23.4/generate-groups.sh

~/go/pkg/mod/k8s.io/code-generator@v0.23.4/generate-groups.sh all "internal/client" "$API_PACKAGE" "$PACKAGE_VERSION" --go-header-file hack/custom-boilerplate.go.txt

mv -f $API_PACKAGE/${PACKAGE_VERSION/://}/*.* internal/api/${PACKAGE_VERSION/://}/
rm -rf github.com

```