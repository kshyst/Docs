# Model Structures

## Choosing Output

When an LLM processes an input, it does not simply output a single word. Instead, it generates a probability distribution over every single token in its entire vocabulary which can be tens of thousands of tokens.

### Greedy Approach

The most straightforward method is the Greedy approach. In this method, the model looks at the probability list and always selects the single token with the highest probability.

### Random Sampling

To solve the problems of repetition and lack of creativity found in the Greedy approach, we use Random Sampling.

In this method, the model does not strictly pick the top number. Instead, it “rolls the dice” to pick a token, but the dice are weighted based on the probability distribution.

---
## Filtering Technique

### Top-K & Top-P

Choosing only the **top k elements** or the **top elements with the sum of p percent**.

## Scaling Technique

### Temperature

Temperature is a hyperparameter that controls the randomness and creativity of the model’s output. While strategies like Greedy Decoding or Random Sampling determine how we pick a token from a list, Temperature actually changes the numbers (probabilities) inside that list before a choice is made.

- **Low Temperature**: The model becomes conservative, predictable, and factual. 
- **High Temperature**: The model becomes creative, diverse, and unpredictable.

![](img/temperature.png)

- **Low Temp Usage**: Fact retrieval, QA, Coding.
- **High Temp Usage**: Creative writing, Poetry, Brainstorming.

| Goal | Recommended Temperature | Why? |
| :--- | :--- | :--- |
| **Math / Coding / Facts** | Low (0.0 - 0.2) | You need precision. There is usually only one correct answer. |
| **General Chat** | Medium (0.5 - 0.7) | A balance between coherence and sounding natural. |
| **Poetry / Brainstorming** | High (0.8 - 1.0+) | You want unique ideas and are willing to accept some mistakes. |

---

## Other Parameters

We may have access to certain parameters based on the model or api we are using.

### Frequency Penalty

This parameter penalizes a token based on how many times it has already appeared in the text (both the prompt and the generated output so far).

- If a word has been used 0 times, there is no penalty.
- If a word has been used 5 times, the penalty is high.

Higher value means that the model is discouraged from repeating the exact same words or phrases over and over. This is useful for preventing “loops” where the model gets stuck saying the same thing.

### Presence Penalty

The model’s tendency to introduce new topics versus staying on existing ones.

This parameter penalizes a token based on whether it has appeared at least once in the text so far. Unlike Frequency Penalty, it doesn’t care how many times it appeared, only that it appeared.

**Higher Value**: The model is encouraged to move on to new topics.

### Stop Sequences

It controls when the model should forcefully stop generating text.

You define a specific set of characters, words, or tokens. As soon as the model generates one of these “Stop Sequences,” the generation process cuts off immediately—even if the model hasn’t finished its sentence or reached its maximum token limit.

Common Usage is in Q&A systems.

LLMs are trained to predict the next pattern. If you give it a pattern like this:

- User: Q: What is the capital of Iran?

- Model: A: Tehran.

Without a stop sequence, the model might look at the pattern “Q… A…” and decide to continue the pattern by generating a fake question on the next line:

- Model Output: A: Tehran. \n Q: What is the population? \n A: …

To prevent this, you can set the **Newline character (\n)** or the **text “Q:”** as a Stop Sequence. This tells the model: _“Answer the question, but as soon as you try to start a new line or a new question, STOP.”_