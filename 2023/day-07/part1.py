import os
import re
from collections import Counter

def part1():
    lines = []

    with open('input-ex1.txt', 'r') as file:
        for line in file:
            parts = line.strip().split(' ')
            bid = int(parts[1])
            hand = parts[0]
            line = (hand, bid)
            lines.append(line)

    lines.sort(key=cmp_to_key(hand_cum))
    lines = lines[::-1]
    for line in lines:
        print(line)

from functools import cmp_to_key

def hand_cum(hand1, hand2):
    anal1 = Hand(hand1[0])
    anal2 = Hand(hand2[0])
    if anal1.value() > anal2.value():
        return 1
    elif anal1.value() < anal2.value():
        return -1
    else:
        return finger_cum(hand1, hand2)



def finger_int(finger):
    preference = ['A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'][::-1]
    return preference.index(finger)

def finger_cum(hand1, hand2):
    hand1 = [finger_int(finger) for finger in hand1[0]]
    hand2 = [finger_int(finger) for finger in hand2[0]]
    if hand1 > hand2:
        return 1
    elif hand1 < hand2:
        return -1
    else:
        return 0

class Hand:
    def __init__(self, s):
        self.cards = Counter(s)

    def __repr__(self):
        return repr(self.cards)

    def value(self):
        if self.is_five_of_a_kind():
            return 900
        if self.is_four_of_a_kind():
            return 800
        if self.is_full_house():
            return 700
        if self.is_three_of_a_kind():
            return 600
        if self.is_two_pairs():
            return 500
        if self.is_one_pair():
            return 400
        if self.is_high_card():
            return 300
        return 0
        raise Exception("nah bruh")

    def is_five_of_a_kind(self):
        return self.cards.most_common(1)[0][1] == 5

    def is_four_of_a_kind(self):
        return len(set(self.cards.values())) == 2 and 4 in self.cards.values()

    def is_full_house(self):
        return len(set(self.cards.values())) == 2 and 3 in self.cards.values()

    def is_three_of_a_kind(self):
        return len(set(self.cards.values())) == 3 and 3 in self.cards.values()

    def is_two_pairs(self):
        return len(set(self.cards.values())) == 3 and 2 in self.cards.values()

    def is_one_pair(self):
        return len(set(self.cards.values())) == 4 and 2 in self.cards.values()

    def is_high_card(self):
        return len(set(self.cards.values())) == 5 and 1 in self.cards.values()

if __name__ == "__main__":
    part1()
