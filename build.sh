if [ ! -d build ]
then
    mkdir ./build   
fi
go build
cp *.js *.html *.css ./build
mv goserve ./build