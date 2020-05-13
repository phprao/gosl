### file generated by Gosl #################################################
import numpy as np
import matplotlib.pyplot as plt
import matplotlib.ticker as tck
import matplotlib.patches as pat
import matplotlib.path as pth
import matplotlib.patheffects as pff
import matplotlib.lines as lns
import mpl_toolkits.mplot3d as m3d
NaN = np.NaN
EXTRA_ARTISTS = []
def addToEA(obj):
    if obj!=None: EXTRA_ARTISTS.append(obj)
COLORMAPS = [plt.cm.bwr, plt.cm.RdBu, plt.cm.hsv, plt.cm.jet, plt.cm.terrain, plt.cm.pink, plt.cm.Greys]
def getCmap(idx): return COLORMAPS[idx % len(COLORMAPS)]
plt.show()
