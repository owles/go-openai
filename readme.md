# Go OpenAI (beta)

The Go OpenAI API module allows to easily interact with the OpenAI API. This module provides a simple and intuitive way to integrate OpenAI's powerful natural language processing capabilities into your Go applications.

[OpenAI Api Documentation](https://platform.openai.com/docs/introduction)

### Features
- Easy to use: The module provides a simple and intuitive interface for working with the OpenAI API. Developers can easily interact with the API without having to worry about the complexities of HTTP requests and JSON parsing.
- Powerful NLP capabilities: The OpenAI API provides a range of powerful natural language processing capabilities, including language translation, text completion, sentiment analysis, and more. This module allows you to harness these capabilities in your Go applications.

### Getting started

To get started with the Go OpenAI API module, you will need an OpenAI API key. You can obtain an API key by creating an account on the OpenAI website.

```
go get github.com/owles/go-openai
```

### Usage

```go

ai, err := openai.NewClient(context.Background(), "YOUR-API-KEY")

if err != nil {
	log.Fatal(err)
}
```

#### Create completion

```go
if c, err := ai.Completion(openai.CompletionRequest{
    Model:  "gpt-3.5-turbo",
    Prompt: "Say this is a test",
}); err == nil {
    for i, choice := range c.Choices {
        fmt.Printf("Completion choice (%d): %s \n", i, choice.Text)
    }
} else {
    log.Fatal(err)
}
```

#### Create chat completion

```go
if c, err := ai.ChatCompletion(openai.ChatCompletionRequest{
    Model: "gpt-3.5-turbo",
    Messages: []openai.Message{
        {
            Role:    openai.RoleUser,
            Content: "Say this is a test by spanish lang!",
        },
    },
}); err == nil {
    for i, choice := range c.Choices {
        fmt.Printf("Chat Completion choice (%d): %s \n", i, choice.Message)
    }
} else {
    log.Fatal(err)
}
```

#### Audio translation

```go
file, _ := os.Open("test.mp3")
defer file.Close()

if tr, err := ai.AudioTranslation(openai.TranslationRequest{
    File:  file,
    Model: "whisper-1",
}); err == nil {
    fmt.Printf("Translation: %s \n", tr.Text)
}
```

#### Generate image

```go
if im, err := ai.GenerateImage(openai.GenerateImageRequest{
    Prompt: "A cute baby sea otter",
    Size:   openai.ImageSize256,
}); err == nil {
    for _, dt := range im.Data {
        fmt.Println(dt.Url)
    }
}
```

