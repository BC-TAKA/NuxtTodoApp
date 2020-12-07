<template>
  <section class="container">
    <v-app>
      <v-main>
        <h3>新規登録欄</h3>
        <v-form>
          <v-container>
            <v-row>
              <v-col cols="12" sm="4">
                Name
                <v-text-field
                label="件名を入力"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="6">
                Todo
                <v-text-field
                label="内容を入力"
                >
                  <template v-slot:append-outer>
                    <v-btn
                      dark
                      rounded
                      color="red"
                    >新規登録</v-btn>
                  </template>
                </v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-form>
        <hr>
        <h3>内容更新欄</h3>
        <v-form>
          <v-container>
            <v-row>
              <v-col cols="12" sm="1">
                ID
                <v-text-field
                  label="IDを入力"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="4">
                Name
                <v-text-field
                  label="件名を入力"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="7">
                Todo
                <v-text-field
                label="内容を入力"
                >
                  <template v-slot:append-outer>
                    <v-btn
                      dark
                      rounded
                      color="blue"
                    >内容を更新</v-btn>
                  </template>
                </v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-form>
        <hr>
        <v-simple-table fixed-header height="500px">
          <template v-slot:default>
            <thead>
              <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Todo</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>No.1</td>
                <td>sample</td>
                <td>samplecase</td>
                <td>
                  <v-btn
                    dark
                    rounded 
                  >削除</v-btn>
                </td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-main>
    </v-app>
  </section>
</template>

<script>
export default {
  data: function() {
    return {
      todos: [],
      createTodoForm: {},
      updateTodoForm: {},
    };
  },
}
</script>