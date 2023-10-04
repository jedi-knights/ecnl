#!/usr/bin/env sh

rm -rf output
mkdir output

start=$(date +%s)
go run main.go rpi --age=G2006/2005 > output/rpi-g2006-2005.txt
end=$(date +%s)
echo "Elapsed Time: $(($end-$start)) seconds"

start=$(date +%s)
go run main.go rpi --age=G2007 > output/rpi-g2007.txt
end=$(date +%s)
echo "Elapsed Time: $(($end-$start)) seconds"

start=$(date +%s)
go run main.go rpi --age=G2008 > output/rpi-g2008.txt
end=$(date +%s)
echo "Elapsed Time: $(($end-$start)) seconds"

start=$(date +%s)
go run main.go rpi --age=G2009 > output/rpi-g2009.txt
end=$(date +%s)
echo "Elapsed Time: $(($end-$start)) seconds"

start=$(date +%s)
go run main.go rpi --age=G2010 > output/rpi-g2010.txt
end=$(date +%s)
echo "Elapsed Time: $(($end-$start)) seconds"

start=$(date +%s)
go run main.go rpi --age=B2006/2005 > output/rpi-b2006-2005.txt
end=$(date +%s)
echo "Elapsed Time: $(($end-$start)) seconds"

start=$(date +%s)
go run main.go rpi --age=B2007 > output/rpi-b2007.txt
end=$(date +%s)
echo "Elapsed Time: $(($end-$start)) seconds"

start=$(date +%s)
go run main.go rpi --age=B2008 > output/rpi-b2008.txt
end=$(date +%s)
echo "Elapsed Time: $(($end-$start)) seconds"

start=$(date +%s)
go run main.go rpi --age=B2009 > output/rpi-b2009.txt
end=$(date +%s)
echo "Elapsed Time: $(($end-$start)) seconds"

start=$(date +%s)
go run main.go rpi --age=B2010 > output/rpi-b2010.txt
end=$(date +%s)
echo "Elapsed Time: $(($end-$start)) seconds"
