# Customizing LLMs

We can customize LLMs to get more precise on certain subjects

> Pre-Trained: Being trained on large general data and being ready to be trained on certain smaller subjects. 

## In-Context Learning

Using certain prompting types to make LLM give the needed answer

Also known as **Few-show Learning** and **Prompt Design**.

## RLHF (Reinforcement Learning from Human Feedback)

Using a **Reward Model** which can be trained to check if the user will accept or decline the main LLM's output. It will use user responses to choose best answers.

This approach is used in ChatGPT

## Fine-Tuning

A common approach in deep learning is to re-adjust a model’s pre-trained parameters for a specific problem. To do this, the
model’s generated output is compared against our desired output. Based on this comparison, the model’s parameters are adjusted (fine-tuned) to improve performance and produce the desired results.

The fine-tuning approach can result in a customized LLM with powerful performance for a specific application. However, it is important to note that this method requires labeled data—meaning we must know the desired output.

> In both Fine-Tuning and RLHF the model's parameters are actively modified. 

