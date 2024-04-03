/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed BY
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/

//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: chainer.proto

package io.seldon.mlops.chainer;

@kotlin.jvm.JvmName("-initializepipelineUpdateStatusMessage")
public inline fun pipelineUpdateStatusMessage(block: io.seldon.mlops.chainer.PipelineUpdateStatusMessageKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusMessage =
  io.seldon.mlops.chainer.PipelineUpdateStatusMessageKt.Dsl._create(io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusMessage.newBuilder()).apply { block() }._build()
public object PipelineUpdateStatusMessageKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusMessage.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusMessage.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusMessage = _builder.build()

    /**
     * <pre>
     * TODO - include `name` to identify transformer message comes from
     * </pre>
     *
     * <code>.seldon.mlops.chainer.PipelineUpdateMessage update = 1;</code>
     */
    public var update: io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateMessage
      @JvmName("getUpdate")
      get() = _builder.getUpdate()
      @JvmName("setUpdate")
      set(value) {
        _builder.setUpdate(value)
      }
    /**
     * <pre>
     * TODO - include `name` to identify transformer message comes from
     * </pre>
     *
     * <code>.seldon.mlops.chainer.PipelineUpdateMessage update = 1;</code>
     */
    public fun clearUpdate() {
      _builder.clearUpdate()
    }
    /**
     * <pre>
     * TODO - include `name` to identify transformer message comes from
     * </pre>
     *
     * <code>.seldon.mlops.chainer.PipelineUpdateMessage update = 1;</code>
     * @return Whether the update field is set.
     */
    public fun hasUpdate(): kotlin.Boolean {
      return _builder.hasUpdate()
    }

    /**
     * <code>bool success = 2;</code>
     */
    public var success: kotlin.Boolean
      @JvmName("getSuccess")
      get() = _builder.getSuccess()
      @JvmName("setSuccess")
      set(value) {
        _builder.setSuccess(value)
      }
    /**
     * <code>bool success = 2;</code>
     */
    public fun clearSuccess() {
      _builder.clearSuccess()
    }

    /**
     * <code>string reason = 3;</code>
     */
    public var reason: kotlin.String
      @JvmName("getReason")
      get() = _builder.getReason()
      @JvmName("setReason")
      set(value) {
        _builder.setReason(value)
      }
    /**
     * <code>string reason = 3;</code>
     */
    public fun clearReason() {
      _builder.clearReason()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusMessage.copy(block: io.seldon.mlops.chainer.PipelineUpdateStatusMessageKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusMessage =
  io.seldon.mlops.chainer.PipelineUpdateStatusMessageKt.Dsl._create(this.toBuilder()).apply { block() }._build()

val io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusMessageOrBuilder.updateOrNull: io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateMessage?
  get() = if (hasUpdate()) getUpdate() else null

