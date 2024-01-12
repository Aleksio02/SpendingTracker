import UsernameTextField from "@/components/account/UsernameTextField.vue";
import PasswordTextField from "@/components/account/PasswordTextField.vue";
import EmailTextField from "@/components/account/EmailTextField.vue";
import FormBlock from "@/components/account/FormBlock.vue";
import FullNameTextField from "@/components/account/FullNameTextField.vue";
import BirthDateTextField from "@/components/account/BirthDateTextField.vue";
import CountryTextField from "@/components/account/CountryTextField.vue";
import CityTextField from "@/components/account/CityTextField.vue";

import SideBar from "@/components/main/SideBar.vue";
import SideBarMenuItem from "@/components/main/SideBarMenuItem.vue";
import PageTitle from "@/components/main/PageTitle.vue";
import PageContent from "@/components/main/PageContent.vue";
import ContentBlockTitle from "@/components/main/ContentBlockTitle.vue";
import ContentBlockText from "@/components/main/ContentBlockText.vue";

import BarChart from "@/components/main/BarChart.vue";
import CalendarBlock from "@/components/main/CalendarBlock.vue";

export default [
    FormBlock, UsernameTextField, PasswordTextField, EmailTextField,
    FullNameTextField, BirthDateTextField, CountryTextField, CityTextField,
    SideBar, SideBarMenuItem, 
    PageTitle, PageContent, ContentBlockTitle, ContentBlockText,
    BarChart, CalendarBlock
]