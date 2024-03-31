# China-Area-Info-Server

超高速内存型省市区查询服务。由Go语言编写。旨在为WEB应用提供高性能的地区查询解决方案

[![Release](https://img.shields.io/github/release/muyu66/china-area-info-server.svg?style=flat-square)](https://github.com/muyu66/china-area-info-server/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

## 特点

* **内存型**：数据源在于内存之中，无需磁盘IO开销
* **Graph模型**：数据组织结构基于Graph模型，拥有极高的瞬时查询能力
* **无畏并发**：万级并发，0.000044516秒/Op
* **极小内存占用**：加载全部数据只占用6.5M内存
* **无锁设计**
* **Restful API**
* **Docker部署** // TODO:
* **GRPC / RPC支持** // TODO:

## Contributors

[Thank you](https://github.com/muyu66/china-area-info-server/graphs/contributors) for contributing!

## License

© Zhouyu, 2024

Released under the [MIT License](https://github.com/muyu66/git-to-dailyreport/blob/master/LICENSE)