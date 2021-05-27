echo "Building web-client build"
cd web-client && yarn build;
echo "Finished web-client build"
cd ..;
echo "embeding web-client started"
rice embed;
echo "embeding web-client done"
