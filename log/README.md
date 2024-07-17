core 核心包含编码和输出
sink 输出
encoder 编码
encoder+sink = zapcore.ioCore

钉钉 loki一类接口就需要特定字段结构化的,最好实现core,当然这往往需要自己实现encoder,sink即第三方接口