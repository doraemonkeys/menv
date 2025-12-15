<p align="center">
  <img src="./assets/logo2.png" width="120" alt="menv logo">
</p>

<h1 align="center">✨ menv ✨</h1>

<p align="center">
  <strong>🪟 귀엽고 강력한 Windows 환경 변수 관리 도구</strong>
</p>

<p align="center">
  <a href="https://github.com/doraemonkeys/menv/actions"><img src="https://github.com/doraemonkeys/menv/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://goreportcard.com/report/github.com/doraemonkeys/menv"><img src="https://goreportcard.com/badge/github.com/doraemonkeys/menv" alt="Go Report Card"></a>
  <a href="https://github.com/doraemonkeys/menv/blob/main/LICENSE"><img src="https://img.shields.io/github/license/doraemonkeys/menv" alt="License"></a>
  <a href="https://github.com/doraemonkeys/menv/releases"><img src="https://img.shields.io/github/v/release/doraemonkeys/menv?include_prereleases" alt="Release"></a>
</p>

<p align="center">
  <a href="README.md">English</a> •
  <a href="README-ZH.md">简体中文</a> •
  <a href="README-ZH-TW.md">繁體中文</a> •
  <a href="README-KO.md">한국어</a> •
  <a href="README-JA.md">日本語</a>
</p>

---

## 🚀 기능

- 📋 **목록 및 검색** - 환경 변수를 쉽게 조회하고 검색
- ✏️ **설정 및 삭제** - 명령줄에서 환경 변수 관리
- 📁 **PATH 관리** - PATH 항목 추가, 제거 및 정리
- 🔍 **상태 검사** - 유효하지 않은 경로 자동 찾기 및 수정
- 💾 **백업 및 복원** - 환경 설정을 절대 잃어버리지 않음
- 🛡️ **시스템 및 사용자** - 사용자 및 시스템 변수 모두 지원

## 📦 설치

[GitHub Releases](https://github.com/doraemonkeys/menv/releases)에서 최신 바이너리를 다운로드하고 PATH에 추가하세요.

또는 Go를 통해 설치:

```bash
go install github.com/doraemonkeys/menv@latest
```

또는 클론하여 로컬에서 설치:

```bash
git clone https://github.com/doraemonkeys/menv.git
cd menv
go install
```

## 🎯 빠른 시작

### 👀 환경 변수 보기

```bash
menv -list                  # 모든 사용자 환경 변수 나열
menv -get JAVA_HOME         # 특정 변수의 값 가져오기
menv -search java           # "java"를 포함하는 변수 검색
```

### ✏️ 환경 변수 설정/삭제

```bash
menv GOPATH C:\Go           # 사용자 환경 변수 설정
menv -d GOPATH              # 사용자 환경 변수 삭제
```

### 📁 PATH 관리

```bash
menv -path                  # PATH 보기 (한 줄에 하나의 경로)
menv -add "C:\bin"          # PATH에 경로 추가
menv -rm "C:\bin"           # PATH에서 경로 제거
menv -clean                 # PATH 정리 (중복 제거 + 유효하지 않은 경로 제거)
menv -search java -path     # PATH에서 "java"를 포함하는 경로 검색
```

### 🔍 유효하지 않은 경로 확인

```bash
menv -check                 # PATH에서 유효하지 않은 디렉토리 확인
menv -check -fix            # 유효하지 않은 경로 자동 제거
menv -check -fix -y         # 확인 없이 바로 제거
```

### 💾 백업 및 복원

```bash
menv -backup backup.json    # 환경 변수 백업
menv -restore backup.json   # 환경 변수 복원
menv -export env.sh         # 셸 스크립트로 내보내기
```

### 🛡️ 시스템 환경 변수

위 명령은 기본적으로 **사용자** 환경 변수에서 작동합니다. **시스템** 환경 변수에서 작동하려면 `-sys`를 추가하세요 (관리자 권한 필요):

```bash
menv -list -sys             # 시스템 환경 변수 나열
menv -path -sys             # 시스템 PATH 보기
menv -add "C:\bin" -sys     # 시스템 PATH에 추가
menv -clean -sys            # 시스템 PATH 정리
```

## 📖 일반적인 시나리오

| 시나리오 | 명령 |
|:---------|:-----|
| 👁️ 특정 변수 보기 | `menv -get JAVA_HOME` |
| ➕ PATH에 디렉토리 추가 | `menv -add "C:\tools\bin"` |
| 🧹 중복 및 유효하지 않은 경로 정리 | `menv -clean` |
| 💾 현재 환경 변수 백업 | `menv -backup my-env.json` |
| 🔎 Java 관련 경로 찾기 | `menv -search java -path` |

## 📄 라이선스

[LICENSE](LICENSE)
