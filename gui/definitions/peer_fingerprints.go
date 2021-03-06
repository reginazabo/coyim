package definitions

func init() {
	add(`PeerFingerprints`, &defPeerFingerprints{})
}

type defPeerFingerprints struct{}

func (*defPeerFingerprints) String() string {
	return `<interface>
  <object class="GtkDialog" id="dialog">
    <property name="title" translatable="yes">Peer fingerprints</property>
    <signal name="close" handler="on_close_signal" />
    <child internal-child="vbox">
      <object class="GtkBox" id="box">
        <property name="border-width">10</property>
        <property name="homogeneous">false</property>
        <property name="orientation">GTK_ORIENTATION_VERTICAL</property>
        <child>
          <object class="GtkLabel" id="information"/>
        </child>
        <child>
          <object class="GtkGrid" id="grid">
            <property name="margin-top">15</property>
            <property name="margin-bottom">10</property>
            <property name="margin-start">10</property>
            <property name="margin-end">10</property>
            <property name="row-spacing">12</property>
            <property name="column-spacing">6</property>
          </object>
        </child>
        <child internal-child="action_area">
          <object class="GtkButtonBox" id="button_box">
            <property name="orientation">GTK_ORIENTATION_HORIZONTAL</property>
            <child>
              <object class="GtkButton" id="button_ok">
                <property name="label" translatable="yes">_OK</property>
                <property name="use-underline">True</property>
                <signal name="clicked" handler="on_close_signal" />
              </object>
            </child>
          </object>
        </child>
      </object>
    </child>
  </object>
</interface>
`
}
