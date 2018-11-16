# histogram - Streaming Multidimensional Approximate Histograms in Go

* This package provides multidimensional [Streaming Approximate Histograms](https://vividcortex.com/blog/2013/07/08/streaming-approximate-histograms/) for efficient quantile approximations.
* The histograms in this package are based on the algorithms found in Ben-Haim & Yom-Tov's *A Streaming Parallel Decision Tree Algorithm* ([PDF](http://jmlr.org/papers/volume11/ben-haim10a/ben-haim10a.pdf)).
* Histogram bins do not have a preset size. As values stream into the histogram, bins are dynamically added and merged.
* A maximum bin size is passed as an argument to the constructor methods. A
larger bin size yields more accurate approximations at the cost of increased
memory utilization and performance.

# Test Results
```
go test -run TestSampleData -timeout 10h

DIMENSION 1
BINS 1024
COUNT 10000
MEAN [-0.3694206912266691]
VARIANCE [9753.369296159837]
CDF(MEAN)       0.5071178239171282 0.5066
CDF(MEAN - 2SD) 0.0216 0.0216
CDF(MEAN - SD)  0.16155923019547438 0.1614
CDF(MEAN + SD)  0.8365931225939377 0.8367
CDF(MEAN + 2SD) 0.9768476589596259 0.9769

DIMENSION 2
BINS 1024
COUNT 10000
MEAN [-0.369420691226669 -0.365370256753095]
VARIANCE [9753.369296159843 10093.457250302596]
CDF(MEAN)       0.247054850062568 0.2526
CDF(MEAN - 2SD) 0.00043209104854484214 0.0005
CDF(MEAN - SD)  0.02526702369621365 0.0257
CDF(MEAN + SD)  0.696886733367091 0.7055
CDF(MEAN + 2SD) 0.9538436217921306 0.954

DIMENSION 3
BINS 1024
COUNT 10000
MEAN [-0.36942069122666815 -0.36537025675309626 -0.6813173206697863]
VARIANCE [9753.369296159844 10093.457250302597 9985.983166140191]
CDF(MEAN)       0.12674345774472287 0.1234
CDF(MEAN - 2SD) 1.113266965637008e-05 0.0001
CDF(MEAN - SD)  0.005058417390372782 0.0041
CDF(MEAN + SD)  0.5783299771273926 0.5949
CDF(MEAN + 2SD) 0.9300958277821303 0.9341

DIMENSION 4
BINS 1024
COUNT 10000
MEAN [-0.36942069122666943 -0.36537025675309603 -0.6813173206697887 0.14363324926603885]
VARIANCE [9753.36929615985 10093.457250302587 9985.983166140193 10044.317379570033]
CDF(MEAN)       0.0619379344047849 0.0653
CDF(MEAN - 2SD) 0 0
CDF(MEAN - SD)  0.0007739989922860771 0.0006
CDF(MEAN + SD)  0.48506764323718016 0.4998
CDF(MEAN + 2SD) 0.9130433633913966 0.9127

DIMENSION 5
BINS 1024
COUNT 10000
MEAN [-0.36942069122666976 -0.3653702567530977 -0.6813173206697868 0.14363324926603896 0.6444562149671539]
VARIANCE [9753.369296159848 10093.457250302588 9985.983166140186 10044.317379570037 10276.776979857257]
CDF(MEAN)       0.03072308838815416 0.0336
CDF(MEAN - 2SD) 0 0
CDF(MEAN - SD)  7.720621790442544e-05 0.0002
CDF(MEAN + SD)  0.4002918386884632 0.4201
CDF(MEAN + 2SD) 0.8878920403025984 0.8915

ok 2229.558s
```

# Comparing CDF with Python
```python
from scipy.stats import mvn
import numpy as np

mu = np.array([0.0, 0.0])
S = np.array([[1.0,0.0],[0.0,1.0]])
low = np.array([-1000.0, -1000.0])

mvn.mvnun(low,np.array([-2.0, -2.0]),mu,S)[0] # CDF(MEAN - 2SD)
mvn.mvnun(low,np.array([-1.0, -1.0]),mu,S)[0] # CDF(MEAN - SD)
mvn.mvnun(low,np.array([0.0, 0.0]),mu,S)[0]   # CDF(MEAN)
mvn.mvnun(low,np.array([1.0, 1.0]),mu,S)[0]   # CDF(MEAN + SD)
mvn.mvnun(low,np.array([2.0, 2.0]),mu,S)[0]   # CDF(MEAN + 2SD)
```

| Dimension |                | Algo                  | Calculated | SciPy                 |
|-----------|----------------|-----------------------|------------|-----------------------|
| 1         | CDF MEAN - 2SD | 0.0216                | 0.0216     | 0.02275013194817923   |
|           | CDF MEAN - SD  | 0.16155923019547438   | 0.1614     | 0.15865525393145702   |
|           | CDF MEAN       | 0.5071178239171282    | 0.5066     | 0.5                   |
|           | CDF MEAN + SD  | 0.8365931225939377    | 0.8367     | 0.841344746068543     |
|           | CDF MEAN + 2SD | 0.9768476589596259    | 0.9769     | 0.9772498680518208    |
| 2         | CDF MEAN - 2SD | 0.00043209104854484214| 0.0005     | 0.0005175685036595823 |
|           | CDF MEAN - SD  | 0.02526702369621365   | 0.0257     | 0.025171489600055108  |
|           | CDF MEAN       | 0.247054850062568     | 0.2526     | 0.25                  |
|           | CDF MEAN + SD  | 0.696886733367091     | 0.7055     | 0.7078609817371412    |
|           | CDF MEAN + 2SD | 0.9538436217921306    | 0.954      | 0.9550173046073012    |
| 3         | CDF MEAN - 2SD | 1.113266965637008e-05 | 0.0001     | 1.1774751750476794e-05|
|           | CDF MEAN - SD  | 0.005058417390372782  | 0.0041     | 0.003993589074329773  |
|           | CDF MEAN       | 0.12674345774472287   | 0.1234     | 0.125                 |
|           | CDF MEAN + SD  | 0.5783299771273926    | 0.5949     | 0.5955551179314647    |
|           | CDF MEAN + 2SD | 0.9300958277821303    | 0.9341     | 0.9332905349146906    |
| 4         | CDF MEAN - 2SD | 0                     | 0          | 2.678771559804014e-07 |
|           | CDF MEAN - SD  | 0.0007739989922860771 | 0.0006     | 0.0006336038886856825 |
|           | CDF MEAN       | 0.0619379344047849    | 0.0653     | 0.625                 |
|           | CDF MEAN + SD  | 0.48506764323718016   | 0.4998     | 0.5010671694658694    |
|           | CDF MEAN + 2SD | 0.9130433633913966    | 0.9127     | 0.9120580520993946    |
| 5         | CDF MEAN - 2SD | 0                     | 0.0216     | 6.09424064445712e-09  |
|           | CDF MEAN - SD  | 7.720621790442544e-05 | 0.0002     | 0.00010052458585138558|
|           | CDF MEAN       | 0.03072308838815416   | 0.0336     | 0.03125               |
|           | CDF MEAN + SD  | 0.4002918386884632    | 0.4201     | 0.4215702304575455    |
|           | CDF MEAN + 2SD | 0.8878920403025984    | 0.8915     | 0.891308611069734     |


# License
MIT License

Copyright (c) 2018 S Sajith

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
