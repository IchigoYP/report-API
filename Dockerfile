# ベースイメージとして最新のUbuntuを使用
FROM ubuntu:latest

# 必要なツールをインストール
RUN apt-get update && apt-get install -y \
    wget \
    git \
    curl \
    vim \
    build-essential

# Go言語をインストール
RUN wget https://dl.google.com/go/go1.16.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz

# Goの環境変数を設定
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/go

# 作業ディレクトリを作成
WORKDIR /app

# ローカルのGoアプリケーションコードをコンテナにコピー
COPY . .

# Goアプリケーションをビルド
RUN go mod init myapp
RUN go build -o myapp .

# コンテナ起動時に実行するコマンド
CMD ["./myapp"]
