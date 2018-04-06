#!/bin/bash
# URL calls for API testing
##### User ADD
# Add user Homero
curl -i -X POST \
--url 'http://localhost:5001/users?user_name=Homero&user_login=Homero&pwd=homerpaswd&email=homero%40whoper.com&gender=machine&location=(0,0)'

# Add user Marge
curl -i -X POST \
--url 'http://localhost:8081/auth/users?user_name=Marge&user_login=Marge&pwd=margepaswd&email=marge%40whoper.com&gender=machine&location=(0,0)'

# Add user Bart
curl -i -X POST \
--url 'http://localhost:8081/auth/users?user_name=Bart&user_login=Bart&pwd=bartpaswd&email=bart%40whoper.com&gender=machine&location=(0,0)'

# Add user Lisa
curl -i -X POST \
--url 'http://localhost:8081/auth/users?user_name=Lisa&user_login=Lisa&pwd=lisapaswd&email=lisa%40whoper.com&gender=machine&location=(0,0)'

# Add user Magie
curl -i -X POST \
--url 'http://localhost:8081/auth/users?user_name=Magie&user_login=Magie&pwd=magiepaswd&email=magie%40whoper.com&gender=machine&location=(0,0)'

# Add user Milhouse
curl -i -X POST \
--url 'http://localhost:8081/auth/users?user_name=Milhouse&user_login=Milhouse&pwd=milhousepaswd&email=milhouse%40whoper.com&gender=machine&location=(0,0)'
##### Login
# Login user Homero
curl -i -X POST \
--url 'http://localhost:8081/auth/validate-username?user_login=Homero&pwd=homerpaswd'
# {"id":1,"user_login":"Homero","salt":"+1FItQtVPXA/okhMMOg2XQtB3/PC+BIsCCt3bXxQIOQQ18AOPA/P8MCIosP/sF3M","token":"YoXr17614d9zKgF89v7MlDiPFF8hnXfw"}

# Login user Marge
curl -i -X POST \
--url 'http://localhost:8081/auth/validate-username?user_login=Marge&pwd=margepaswd'
# {"id":2,"user_login":"Marge","salt":"Xb823xBBMPXhboDkaoQXkqaAxFQIjbKF+XQhCQA8/XIX+IC2M3CQgMOxQIxBIH/o","token":"yKU6S4cxPITUWfls9xOwXsANlwXde7nu"}

# Login user Bart
curl -i -X POST \
--url 'http://localhost:8081/auth/validate-username?user_login=Bart&pwd=bartpaswd'
# {"id":3,"user_login":"Bart","salt":"3+oM32CAHBxFIooxkADA+HBCOaQhtMBoI/C11g88oQAQIh1A1BQhQxgBXsqB/8Xb","token":"MZFMhpfOJBQjVgeTBAaarESFoW38RtKr"}

# Login user Lisa
curl -i -X POST \
--url 'http://localhost:8081/auth/validate-username?user_login=Lisa&pwd=lisapaswd'
# {"id":4,"user_login":"Lisa","salt":"hIPKkM3o3jPCCaXb3Q/BAskIqPoXoXC3PqkBQ8o8jPDMMB2CK8C8s2/BBM81BXqP","token":"hovn2pyuWvbGu6oef0E4E92RKW0a7TmU"}

# Login user Magie
curl -i -X POST \
--url 'http://localhost:8081/auth/validate-username?user_login=Magie&pwd=magiepaswd'
# {"id":5,"user_login":"Magie","salt":"VDo8IbsMVPXVDDIocBBoQoCeCQIXQ2IQADQIg//o+3CCXCABMBBsx2kK/aBIoQkk","token":"fDYorXMn3HBYOOrKIOYy4FwHoFjiSRcD"}

# login user Milhouse
curl -i -X POST \
> --url 'http://localhost:8081/auth/validate-username?user_login=Milhouse&pwd=milhousepaswd'
# {"id":7,"user_login":"Milhouse","salt":"PkQk13+bkBK11CBQQ8VO1oIXoXxhMQBOMQ+PaKs1A3A/q1aqb/h8HXQQooBQQsXC","token":"6qt93L9Wjjwg0f36anTtHXORZzL68IhP"}

# Social network
####  Homer
# Homer Follows Marge
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/1/2' -H 'apikey: YoXr17614d9zKgF89v7MlDiPFF8hnXfw'

# Homer Follows Bart
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/1/3' -H 'apikey: YoXr17614d9zKgF89v7MlDiPFF8hnXfw'

# Homer Follows Lisa
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/1/4' -H 'apikey: YoXr17614d9zKgF89v7MlDiPFF8hnXfw'

# Homer Follows Magie
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/1/5' -H 'apikey: YoXr17614d9zKgF89v7MlDiPFF8hnXfw'
####  Marge
# Marge Follows Homer
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/2/1' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Marge Follows Bart
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/2/3' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Marge Follows Lisa
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/2/4' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Marge Follows Magie
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/2/5' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'
####  Bart
# Bart Follows Homer
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/3/1' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Bart Follows Marge
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/3/2' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Bart Follows Lisa
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/3/4' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Bart Follows Magie
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/3/5' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Bart Follows Milhouse
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/3/7' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'
####  Lisa
# Lisa Follows Homer
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/4/1' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Lisa Follows Marge
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/4/2' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Lisa Follows Bart
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/4/3' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Lisa Follows Magie
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/4/5' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'
####  Magie
# Magie Follows Homer
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/5/1' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Magie Follows Marge
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/5/2' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Magie Follows Lisa
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/5/3' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

# Magie Follows Bart
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/5/4' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

####  Milhouse
# Milhouse Follows Bart
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/ntw/7/3' -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'
##### Posts
# Posts of Homero
# Post No-1 Public

curl -i -X POST \
--url 'http://localhost:8081/postsapi/posts?price=10000&post_owner=1&rating=4&location=%281%2C1%29&caption=blue%20pants&group=Public' \
-H 'apikey: YoXr17614d9zKgF89v7MlDiPFF8hnXfw'
# {"id":4,"post_date":"2017-09-28T00:00:00Z","post_rating":"4","price":10000,"location":"(1,1)","caption":"blue pants","status":1,"owner_id":1}

curl -i -X POST \
--url 'http://localhost:8081/imagesapi/image' \
--data 'url=https://static2.fjcdn.com/thumbnails/comments/5488497+_35ece9585d73add9dc2d0ab5489a0198.png' \
--data 'post_id=4' \
-H 'apikey: YoXr17614d9zKgF89v7MlDiPFF8hnXfw'

# Post No-1 Private
# Create a private group
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/groups' \
--data 'owner_id=1' \
--data 'group_name=forMarge' \
 -H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'
# Add users to the group
curl -i -X POST \
--url 'http://localhost:8081/ntwrsapi/groups/1/11/2' \
-H 'apikey: yKU6S4cxPITUWfls9xOwXsANlwXde7nu'

curl -i -X POST \
--url 'http://localhost:8081/postsapi/posts?price=10000&post_owner=1&rating=4&location=%281%2C1%29&caption=blue%20pants&group=Private' \
-H 'apikey: YoXr17614d9zKgF89v7MlDiPFF8hnXfw'
# {"id":4,"post_date":"2017-09-28T00:00:00Z","post_rating":"4","price":10000,"location":"(1,1)","caption":"blue pants","status":1,"owner_id":1}

curl -i -X POST \
--url 'http://localhost:8081/imagesapi/image' \
--data 'url=https://static2.fjcdn.com/thumbnails/comments/5488497+_35ece9585d73add9dc2d0ab5489a0198.png' \
-H 'apikey: YoXr17614d9zKgF89v7MlDiPFF8hnXfw'
# EOF!
