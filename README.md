# simple-website-cli
A simple website CLI to query NS, IP, CNAME, MX, and TXT

Inspired by the this tutorial: https://tutorialedge.net/golang/building-a-cli-in-go/


Build CLI: ```go build path/net-cli.go``` 

Copy: ```cp net-cli /usr/local/bin```

Name Servers: ```net-cli ns --host example.com```

IP: ```net-cli ip --host example.com```

CNAME: ```net-cli cname --host example.com```

MX Records: ```net-cli mx --host example.com```

TXT Records: ```net-cli txt --host example.com```
