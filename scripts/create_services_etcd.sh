curl -XDELETE http://10.89.4.165:4001/v2/keys/services?recursive=true&dir=true

curl -XPOST http://10.89.4.165:4001/v2/keys/services/cad/domain -d value=cad
curl -XPOST http://10.89.4.165:4001/v2/keys/services/admin/domain -d value=admin

curl -XPOST http://10.89.4.165:4001/v2/keys/services/cad/backend/one -d value=10.89.4.165:3000
curl -XPOST http://10.89.4.165:4001/v2/keys/services/admin/backend/two -d value=10.89.4.165:3001

curl -XPOST http://10.89.4.165:4001/v2/keys/services/cad/backend/three -d value=10.89.4.165:3002
curl -XPOST http://10.89.4.165:4001/v2/keys/services/admin/backend/four -d value=10.89.4.165:3003

curl -XPOST http://10.89.4.165:4001/v2/keys/services/cad/backend/five -d value=10.89.4.165:3004
