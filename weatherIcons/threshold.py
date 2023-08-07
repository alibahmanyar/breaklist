#! /usr/bin/python3
import numpy as np, time, cv2
from matplotlib import pyplot as plt
import os

files = os.listdir("./raw")
WIN_SIZE = 11
C = 3

for f in files:
    a = cv2.imread("./raw/" + f)

    rt = False

    if a.shape[1] > a.shape[0]:
        rt = True
        a = cv2.rotate(a, cv2.cv2.ROTATE_90_CLOCKWISE) 

    print(a.shape)

    a = cv2.cvtColor(a, cv2.COLOR_BGR2GRAY)

    a = cv2.adaptiveThreshold(a,255,cv2.ADAPTIVE_THRESH_GAUSSIAN_C,\
                cv2.THRESH_BINARY, WIN_SIZE, C)
    
    # res, a = cv2.threshold(a,127,255,cv2.THRESH_BINARY)

    a = cv2.bitwise_not(a)

    a = cv2.resize(a, (384, int(a.shape[0] * (384/a.shape[1]))))

    b = a
    if rt:
        b = cv2.rotate(b, cv2.cv2.ROTATE_90_COUNTERCLOCKWISE)

    cv2.imwrite("./thresholded/"+f, cv2.bitwise_not(b))
    # plt.imshow(b, cmap='Greys')
    # plt.show()
