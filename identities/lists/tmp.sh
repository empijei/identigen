#todo iterate over page numbers too!
for lettera in Q W E R T Y U I O P A S D F G H J K L Z X C V B N M
do
wget "http://www.cognomix.it/origine-cognomi-italiani/$lettera"
grep -Po 'php" title="\K[^-]+' $lettera
done > cognomix.txt
