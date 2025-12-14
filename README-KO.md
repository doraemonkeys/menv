# menv

[English](README.md) | [简体中文](README-ZH.md) | [繁體中文](README-ZH-TW.md) | [한국어](README-KO.md) | [日本語](README-JA.md)

명령줄에서 환경 변수와 PATH를 쉽게 관리할 수 있는 Windows 환경 변수 관리 도구입니다.

## 설치

```bash
go install github.com/doraemonkeys/menv@latest
```

또는 클론 후 로컬 설치:

```bash
git clone https://github.com/doraemonkeys/menv.git
cd menv
go install
```

## 빠른 시작

### 환경 변수 보기

```bash
menv -list                  # 모든 사용자 환경 변수 나열
menv -get JAVA_HOME         # 특정 변수의 값 가져오기
menv -search java           # java를 포함하는 변수 검색
```

### 환경 변수 설정/삭제

```bash
menv GOPATH C:\Go           # 사용자 환경 변수 설정
menv -d GOPATH              # 사용자 환경 변수 삭제
```

### PATH 관리

```bash
menv -path                  # PATH 보기 (한 줄에 하나의 경로)
menv -add "C:\bin"          # PATH에 경로 추가
menv -rm "C:\bin"           # PATH에서 경로 제거
menv -clean                 # PATH 정리 (중복 제거 + 잘못된 경로 삭제)
menv -search java -path     # PATH에서 java를 포함하는 경로 검색
```

### 잘못된 경로 확인

```bash
menv -check                 # PATH에서 잘못된 디렉터리 확인
menv -check -fix            # 잘못된 경로 자동 제거
menv -check -fix -i         # 각 경로 제거 전 확인
```

### 백업 및 복원

```bash
menv -backup backup.json    # 환경 변수 백업
menv -restore backup.json   # 환경 변수 복원
menv -export env.sh         # 셸 스크립트로 내보내기
```

### 시스템 환경 변수

위 명령어는 기본적으로 **사용자** 환경 변수를 조작합니다. **시스템** 환경 변수를 조작하려면 `-sys`를 추가하세요 (관리자 권한 필요):

```bash
menv -list -sys             # 시스템 환경 변수 나열
menv -path -sys             # 시스템 PATH 보기
menv -add "C:\bin" -sys     # 시스템 PATH에 추가
menv -clean -sys            # 시스템 PATH 정리
```

## 일반적인 시나리오

| 시나리오 | 명령어 |
|----------|--------|
| 특정 변수 보기 | `menv -get JAVA_HOME` |
| PATH에 디렉터리 추가 | `menv -add "C:\tools\bin"` |
| 중복 및 잘못된 경로 정리 | `menv -clean` |
| 현재 환경 변수 백업 | `menv -backup my-env.json` |
| Java 관련 경로 찾기 | `menv -search java -path` |
