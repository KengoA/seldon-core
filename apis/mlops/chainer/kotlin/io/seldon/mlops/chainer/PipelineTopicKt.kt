//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: chainer.proto

package io.seldon.mlops.chainer;

@kotlin.jvm.JvmName("-initializepipelineTopic")
public inline fun pipelineTopic(block: io.seldon.mlops.chainer.PipelineTopicKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.chainer.ChainerOuterClass.PipelineTopic =
  io.seldon.mlops.chainer.PipelineTopicKt.Dsl._create(io.seldon.mlops.chainer.ChainerOuterClass.PipelineTopic.newBuilder()).apply { block() }._build()
public object PipelineTopicKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: io.seldon.mlops.chainer.ChainerOuterClass.PipelineTopic.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: io.seldon.mlops.chainer.ChainerOuterClass.PipelineTopic.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): io.seldon.mlops.chainer.ChainerOuterClass.PipelineTopic = _builder.build()

    /**
     * <code>string pipelineName = 1;</code>
     */
    public var pipelineName: kotlin.String
      @JvmName("getPipelineName")
      get() = _builder.getPipelineName()
      @JvmName("setPipelineName")
      set(value) {
        _builder.setPipelineName(value)
      }
    /**
     * <code>string pipelineName = 1;</code>
     */
    public fun clearPipelineName() {
      _builder.clearPipelineName()
    }

    /**
     * <code>int64 pipelineVersion = 2;</code>
     */
    public var pipelineVersion: kotlin.Long
      @JvmName("getPipelineVersion")
      get() = _builder.getPipelineVersion()
      @JvmName("setPipelineVersion")
      set(value) {
        _builder.setPipelineVersion(value)
      }
    /**
     * <code>int64 pipelineVersion = 2;</code>
     */
    public fun clearPipelineVersion() {
      _builder.clearPipelineVersion()
    }

    /**
     * <code>string topicName = 3;</code>
     */
    public var topicName: kotlin.String
      @JvmName("getTopicName")
      get() = _builder.getTopicName()
      @JvmName("setTopicName")
      set(value) {
        _builder.setTopicName(value)
      }
    /**
     * <code>string topicName = 3;</code>
     */
    public fun clearTopicName() {
      _builder.clearTopicName()
    }

    /**
     * <code>optional string tensor = 4;</code>
     */
    public var tensor: kotlin.String
      @JvmName("getTensor")
      get() = _builder.getTensor()
      @JvmName("setTensor")
      set(value) {
        _builder.setTensor(value)
      }
    /**
     * <code>optional string tensor = 4;</code>
     */
    public fun clearTensor() {
      _builder.clearTensor()
    }
    /**
     * <code>optional string tensor = 4;</code>
     * @return Whether the tensor field is set.
     */
    public fun hasTensor(): kotlin.Boolean {
      return _builder.hasTensor()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun io.seldon.mlops.chainer.ChainerOuterClass.PipelineTopic.copy(block: io.seldon.mlops.chainer.PipelineTopicKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.chainer.ChainerOuterClass.PipelineTopic =
  io.seldon.mlops.chainer.PipelineTopicKt.Dsl._create(this.toBuilder()).apply { block() }._build()

