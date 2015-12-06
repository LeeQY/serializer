# serializer
Provide some methods to do the serialization job.

###Install & Update
```
go get -u github.com/LeeQY/serializer
```

###Usage

```Go
// Encode an int array to string seperate by ","
s, err := serializer.EncodeIntArray(intArray)

// Decode a string to int array split by ","
intArray, err := serializer.DecodeIntArray(s)

// Encode a string array to string seperate by strSep
s, err := serializer.EncodeStringArray(strArray, strSep)

// Decode a string to string array split by strSep
strArray, err := serializer.DecodeStringArray(&s, strSep)

```