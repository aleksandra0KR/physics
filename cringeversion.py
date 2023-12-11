import sys
import pygame
import sympy as sp
import pygame_gui


'''
## тут какой-то лютый кринж, хз зачем он, но с ним раболтает все заебись, но заебеннее, поэтому чекнуть, что тут происходит
# Уравнение сохранения импульса
momentum_eq = sp.Eq(
    m1 * v1i + m2 * v2i,
    m1 * v1f + m2 * v2f
)
# Уравнение сохранения кинетической энергии
energy_eq = sp.Eq(
    half * m1 * v1i ** 2 + half * m2 * v2i ** 2,
    half * m1 * v1f ** 2 + half * m2 * v2f ** 2
)
# Решение системы уравнений относительно v1f и v2f
solutions = sp.solve([momentum_eq, energy_eq], [v1f, v2f])

# Отбор решений, которые соответствуют физическим условиям (скорости должны измениться после столкновения)
solutions = [s for s in solutions if s != (v1i, v2i)][0]
v1f_solved, v2f_solved = [sp.simplify(s) for s in solutions]


# Лямбда-функции для вычисления скоростей после столкновения на основе полученных решений
##v1f_f = sp.lambdify([m1, m2, v1i, v2i], v1f_solved)
##v2f_f = sp.lambdify([m1, m2, v1i, v2i], v2f_solved)

# Инициализация pygame

'''


# settings for screen
WIDTH, HEIGHT = 800, 400
screen = pygame.display.set_mode((WIDTH, HEIGHT))
pygame.display.set_caption("Absolutely elastic interaction")
pygame.init()
manager = pygame_gui.UIManager((800, 400))

text_input = pygame_gui.elements.UITextEntryLine(relative_rect=pygame.Rect((230, 200), (400, 50)), manager=manager,
                                                 object_id='#main_text_entry')

m1, m2, v1i, v2i, v1f, v2f = sp.symbols('m_1, m_2, v_1i, v_2i, v_1f, v_2f')
global mass1, mass2, speed2
half = sp.numer(1) / 2


# colors to use // now we don't fully use them
YELLOW = (255, 255, 0)
RED = (255, 0, 0)
BLACK = (0, 0, 0)
WHITE = (255, 255, 255)


# entety for block
class Block:
    def __init__(self, x, y, mass, velocity, color):
        self.mass = mass
        self.velocity = velocity
        self.color = color
        self.size = int(40 * (mass / 10) ** (1 / 3))
        self.rect = pygame.Rect(x, y - self.size, self.size, self.size)

    def move(self):
        self.rect.x += self.velocity
        if self.rect.x > 800:
            pygame.quit()
            sys.exit()

    def collide_with_wall(self):
        if self.rect.x <= 0:
            self.rect.x = 0  # Коррекция положения, чтобы блок не уходил за пределы экрана
            self.velocity = -self.velocity  # Изменение направления движения
            return True
        return False

    def draw(self):
        pygame.draw.rect(screen, self.color, self.rect)
        font = pygame.font.SysFont(None, 25)
        mass_text = font.render(str(self.mass), True, BLACK)
        screen.blit(mass_text, (self.rect.x + self.size // 2 - mass_text.get_width() // 2,
                                self.rect.y + self.size // 2 - mass_text.get_height() // 2))


# counting speed after collision with object_2 for self element
def get_speed_after_collision(self, object_2):
    return ((self.mass - object_2.mass) / (self.mass + object_2.mass)) * self.velocity + \
        ((2 * object_2.mass) / (self.mass + object_2.mass)) * object_2.velocity

# shit
def calculate_collision_velocity(b1, b2):
    return get_speed_after_collision(b1, b2), get_speed_after_collision(b2, b1)


def main(mass1, mass2, speed2):

    # creating blocks
    block1 = Block(200, HEIGHT, mass1, 0, (192, 185, 221))
    block2 = Block(500, HEIGHT, mass2, speed2, (117, 201, 200))

    collision_count = 0

    # main loop
    running = True
    while running:
        screen.fill("white")

        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False

        block1.move()
        block2.move()

        if block1.rect.colliderect(block2.rect):
            v1_final, v2_final = calculate_collision_velocity(block1, block2)
            block1.velocity = v1_final
            block2.velocity = v2_final

            # Минимальное смещение, чтобы избежать зацикливания столкновений
            while block1.rect.colliderect(block2.rect):
                block1.rect.x += block1.velocity
                block2.rect.x += block2.velocity

            collision_count += 1

        if block1.collide_with_wall():
            block1.velocity = abs(block1.velocity)
            collision_count += 1

        block1.draw()
        block2.draw()

        # Отображаем количество столкновений
        font = pygame.font.SysFont(None, 35)
        collision_text = font.render(f"Collisions: {collision_count}", True, BLACK)
        screen.blit(collision_text, (WIDTH // 2 - collision_text.get_width() // 2, 10))

        pygame.display.flip()
        pygame.time.Clock().tick(60)

    pygame.quit()
    sys.exit()


def get_data():
    while True:

        for event in pygame.event.get():

            if event.type == pygame.QUIT:
                pygame.quit()
                sys.exit()
            if (event.type == pygame_gui.UI_TEXT_ENTRY_FINISHED and
                    event.ui_object_id == '#main_text_entry'):
                mass1, mass2, speed2 = [int(x) for x in text_input.get_text().split()]
                main(mass1, mass2, -speed2)
            manager.process_events(event)
        manager.update(10)

        screen.fill("white")
        font = pygame.font.SysFont(None, 24)
        text_surface = font.render("Enter masses of body 1 and body 2 and speed of 2 body separated by a space:", False,
                                   "black")
        text_rect = text_surface.get_rect()
        text_rect.centerx = screen.get_width() / 2
        text_rect.centery = screen.get_height() / 2 - 50
        screen.blit(text_surface, text_rect)

        manager.draw_ui(screen)

        pygame.display.update()

def get_data():
    while True:

        for event in pygame.event.get():

            if event.type == pygame.QUIT:
                pygame.quit()
                sys.exit()
            if (event.type == pygame_gui.UI_TEXT_ENTRY_FINISHED and
                    event.ui_object_id == '#main_text_entry'):
                mass1, mass2, speed2 = [int(x) for x in text_input.get_text().split()]
                main(mass1, mass2, -speed2)
            manager.process_events(event)
        manager.update(10)

        screen.fill("white")
        font = pygame.font.SysFont(None, 24)
        text_surface = font.render("Enter masses of body 1 and body 2 and speed of 2 body separated by a space:", False,
                                   "black")
        text_rect = text_surface.get_rect()
        text_rect.centerx = screen.get_width() / 2
        text_rect.centery = screen.get_height() / 2 - 50
        screen.blit(text_surface, text_rect)

        manager.draw_ui(screen)

        pygame.display.update()



pygame.display.update()
get_data()
