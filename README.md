### urlencoder
Encode and decode url 


Ever found yourself trying to url encode certain characters when doing pentest?

Urlencoder will help you **encode**,**double encode** or **decode**

Encoding a payload will give two results,the encoded version and the double encoded version

### Requirements

Requires golang to be installed and the gopath set 

### Installation

```
go get github.com/peter-macharia/urlencoder

```

### Usage

echo "string_to_encode" | urlencoder

echo "string_to_decode" | urlencoder -d
