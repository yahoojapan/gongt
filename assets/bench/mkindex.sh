#!/bin/sh

./example -create -name Fashion-MNIST -path fashion-mnist-784-euclidean.hdf5
./example -create -name GloVe-25 -path glove-25-angular.hdf5
./example -create -name GloVe-50 -path glove-50-angular.hdf5
./example -create -name GloVe-100 -path glove-100-angular.hdf5
#./example -create -name GloVe-200 -path glove-200-angular.hdf5
./example -create -name MNIST -path mnist-784-euclidean.hdf5
./example -create -name NYTimes -path nytimes-256-angular.hdf5
./example -create -name SIFT -path sift-128-euclidean.hdf5
