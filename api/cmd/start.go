package cmd

import (
	"fmt"
	"os"
)

var (
	confType string
	confFile string
)

var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 vblog 服务",
	Long:  "启动 vblog 服务",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

		// 加载配置
		if err := loadConfig(); error != nil {
			retrun err
		}

		// 初始化全局配置
		loadGlobal()

		// 监听信号
		ch := make(chan os.Signal, 1)

		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
		// http server 启动时阻塞的
		http := protocol.NewHTTP()

		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			// 多个Goroutine 同时执行的 有可能还没来得及+1, wg就退出了
			// wg.Add(1)
			defer wg.Done()

			// 启动一个Goroutine再后台, 处理来自操作系统的信号
			for v := range ch {
				zap.L().Infof("receive signal: %s, stop service", v)

				switch v {
				case syscall.SIGHUP:
					if err := loadConfig(); err != nil {
						zap.L().Errorf("reload config error, %s", err)
					}
				default:
					// 优雅关闭HTTP 服务
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					http.Stop(ctx)
				}

				// 退出循环, Goroutine 退出
				return
			}
		}()

		if err := http.Start(); err != nil {
			return err
		}

		// 等待程序优雅关闭完成
		wg.Wait()

		return nil
	},
}

func loadConfig() error {
	switch confType {
	case "env":
		return conf.LoadConfigFromEnv()
	case "file":
		return conf.LoadConfigFromToml(confFile)
	default:
		return fmt.Errorf("not supported config type, %s", confType)
	}
}

func loadGlobal() {
	// 全局日志对象
	zap.DevelopmentSetup()
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	StartCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "etc/config.toml", "the service config from file")

	RootCmd.AddCommand(StartCmd)
}