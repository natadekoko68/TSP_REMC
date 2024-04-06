path_init = "/Users/kotaro/Desktop/MCMC_2d_init.txt"
path_fin = "/Users/kotaro/Desktop/MCMC_2d.txt"

import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
import mpl_toolkits.mplot3d.art3d as art3d
import numpy as np

cities = []
x = []
y = []
with open(path_fin) as f:
    for line in f:
        temp_x, temp_y = line.rstrip().split()
        x.append(float(temp_x))
        y.append(float(temp_y))
x = np.array(x)
y = np.array(y)

fig = plt.figure()
plt.scatter(x, y, marker="+", color="k")
plt.plot(x, y, color="b")
plt.text(x[0], y[0], "s", color='green')
plt.text(x[-1], y[-1], "g", color='green')

plt.savefig("/Users/kotaro/Desktop/MCMC_2d.png", dpi=500)
plt.show()


