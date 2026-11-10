defmodule InterviewSetup.MixProject do
  use Mix.Project

  def project do
    [
      app: :interview_setup,
      version: "0.1.0",
      elixir: "~> 1.17",
      deps: [{:exqlite, "~> 0.35"}, {:jason, "~> 1.4"}]
    ]
  end

  def application, do: [extra_applications: [:logger]]
end
