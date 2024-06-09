import numpy as np
import matplotlib.pyplot as plt

# Параметры линзы и света
radius = 100  # mm, радиус линзы
wavelength_center = 600  # nm, середина спектра
wavelength_width = 100  # nm, ширина спектра
distance = 1000  # mm, расстояние от линзы до экрана

# Рассчет координат и интенсивности колец Ньютона
radial_coords = np.linspace(0, 2 * radius, 400)  # массив радиальных координат
wavelengths = np.linspace(wavelength_center - wavelength_width / 2,
                            wavelength_center + wavelength_width / 2, 100)  # массив длин волн

intensity_monochromatic = np.zeros_like(radial_coords)
intensity_quasi_monochromatic = np.zeros_like(radial_coords)

for i, r in enumerate(radial_coords):
    # Расчет фазовой разницы между лучами, проходящими через край и центр линзы
    phase_difference = (r ** 2 / (distance * wavelength_center)) + \
                         (r ** 2 / (distance * wavelength_center)) ** 2

    # Расчет интенсивности для монохроматического света
    intensity_monochromatic[i] = (1 + np.cos(2 * np.pi * phase_difference)) ** 2

    # Расчет интенсивности для квазимонохроматического света
    for w in wavelengths:
        phase_difference_w = (r ** 2 / (distance * w)) + \
                               (r ** 2 / (distance * w)) ** 2
        intensity_quasi_monochromatic[i] += (1 + np.cos(2 * np.pi * phase_difference_w)) ** 2

# Визуализация результатов
plt.figure(figsize=(10, 5))

# Цветное распределение интенсивности интерференционной картины
plt.subplot(121)
plt.imshow(np.tile(intensity_quasi_monochromatic[:, np.newaxis], (1, len(radial_coords))),
           extent=[radial_coords.min(), radial_coords.max(), radial_coords.max(), radial_coords.min()],
           cmap='inferno', origin='lower')
plt.xlabel('Radial coordinate, mm')
plt.ylabel('Radial coordinate, mm')

# График зависимости интенсивности от радиальной координаты
plt.subplot(122)
plt.plot(radial_coords, intensity_monochromatic, label='Monochromatic light')
plt.plot(radial_coords, intensity_quasi_monochromatic, label='Quasi-monochromatic light')
plt.xlabel('Radial coordinate, mm')
plt.ylabel('Intensity')
plt.legend()

plt.show()
