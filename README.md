# markov-chain

Markov Chain Algorithm

## Abstract
In this project I built a text generator using *Markov Chain algorithm.*

Similar algorithms are used in your phones. For example, when you type a word in the
keyboard it suggests you the next the most probable word.

## Learning Objectives
    Algorithms
    I/O
    File Handling
    Basic software design principles

Mine specific task is to generate random English text that reads naturally. If we were to emit random letters or random words, the output would be nonsensical. For instance, a program that randomly picks letters (and spaces to separate words) might generate something like this:
    tqrhyla kjsbq oewzj xna i smdkt uivbxfytnvplxq

Similarly, selecting words at random from a dictionary won't produce coherent text either:
    almanac threshold griddle 9f zigzag cactus haphazardly blue anchor

To achieve more meaningful results, we need a statistical model with greater structure, such as one that considers the frequency of entire phrases. But how can we obtain such statistical data?

Markov Chain Algorithm

An intelligent way to handle this type of text generation is with a technique called a Markov chain algorithm. This method views the input as a series of overlapping sequences, breaking each sequence into two parts: a multi-word prefix and a single suffix word that follows the prefix. The algorithm creates new sequences by randomly picking a suffix that follows each prefix, based on patterns in the original text. Using three-word sequences, with a two-word prefix to select the suffix, works well:

1. Set w1 and w2 to the first two words of the text.
2. Print w1 and w2.
3. Loop:
    a. Randomly choose w3, a word that can follow w1 and w2 in the text.
    b. Print w3.
    c. Update w1 to w2 and w2 to w3.
    d. Repeat loop.

## Mandatory Part

### Baseline
Notes:
    Suffix length is ALWAYS 1 word.
    Default prefix length is 2 words.
    Default starting prefix is the first N words of the text, where N is the length of the prefix.
    Default number of maximum words is 100.

Constraints:
    If any error print an error message indicating the reason.
    The code should stop generating code after it printed maximum number of words or encountered the very last word in the text.

Examples:
    $ cat the_great_gatsby.txt | ./markovchain | cat -e
    Chapter 1 In my younger and more stable, become for a job. He hadn't eat anything for a long, silent time. It was the sound of someone splashing after us over the confusion a long many-windowed room which overhung the terrace. Eluding Jordan's undergraduate who was well over sixty, and Maurice A. Flink and the great bursts of leaves growing on the air now. "How do you want? What do you like Europe?" she exclaimed surprisingly. "I just got here a minute. "Yes." He hesitated. "Was she killed?" "Yes." "I thought you didn't, if you'll pardon my--you see, I carry$

    $ cat the_great_gatsby.txt | ./markovchain | wc -w
    100

    $ ./markovchain
    Error: no input text

### Number of words
Outcomes:
    Program prints generated text according to the Markov Chain algorithm limited by the given maximum number of words.

Constraints:
    Given number can't be negative.
    Given number can't be more 10,000.
    If any error print an error message indicating the reason.

Example:
    $ cat the_great_gatsby.txt | ./markovchain -w 10 | cat -e
    Chapter 1 In my younger and more stable, become for$

### Prefix
Outcomes:
    Program prints generated text according to the Markov Chain algorithm that starts with the given prefix.

Constraints:
    Given prefix must be present in the original text.
    If any error print an error message indicating the reason.

Example: 
    $ cat the_great_gatsby.txt | ./markovchain -w 10 -p "to play" | cat -e
    to play for you in that vast obscurity beyond the$

### Prefix length
Outcomes:
    Program prints generated text according to the Markov Chain algorithm with the given prefix length.

Constraints:
    Given prefix length can't be negative.
    Given prefix length can't be greater than 5.
    If any error print an error message indicating the reason.

Example:
    $ cat the_great_gatsby.txt | ./markovchain -w 10 -p "to something funny" -l 3
    to something funny the last two days," remarked Wilson. "That's

### Usage
Outcomes:
    Program prints usage text.
Example:
    $ ./markovchain --help
    Markov Chain text generator.

    Usage:
    markovchain [-w <N>] [-p <S>] [-l <N>]
    markovchain --help

    Options:
    --help  Show this screen.
    -w N    Number of maximum words
    -p S    Starting prefix
    -l N    Prefix length


